from abc import ABC, abstractmethod
from dataclasses import dataclass
from timeit import default_timer
from typing import Optional

import aio_pika
import grpc
import structlog
from aio_pika.abc import AbstractRobustConnection
from nagi.inference import Caption, QuestionAndAnswer
from nagi.storage import Storage

from hskw import hskw_pb2, hskw_pb2_grpc

log: structlog.stdlib.BoundLogger = structlog.get_logger()


@dataclass
class QueueConfig:
    address: str
    queue_name: str
    prefetch_count: int = 1


class AbstractQueue(ABC):
    def __init__(self, config: QueueConfig):
        self.connection = Optional[AbstractRobustConnection]
        self.address = config.address
        self.queue_name = config.queue_name
        self.prefetch_count = config.prefetch_count

    @abstractmethod
    async def process_message(
        self, message: aio_pika.abc.AbstractIncomingMessage
    ) -> None:
        """Process a message from the queue"""

    async def connect(self) -> AbstractRobustConnection:
        return await aio_pika.connect_robust(self.address)

    async def run(self):
        self.connection = await self.connect()
        log.info("Connected to queue", address=self.address)
        channel = await self.connection.channel()
        await channel.set_qos(prefetch_count=self.prefetch_count)

        queue = await channel.declare_queue(self.queue_name, auto_delete=False)
        log.info("Starting to consume", queue_name=self.queue_name)
        await queue.consume(self.process_message)


class CaptionQueue(AbstractQueue):
    def __init__(self, config: QueueConfig, caption_model: Caption, storage: Storage, grpc_endpoint: str):
        super().__init__(config)
        self.model = caption_model
        self.storage = storage
        self.grpc_endpoint = grpc_endpoint

    def parse_message(
        self, message: aio_pika.abc.AbstractIncomingMessage
    ) -> hskw_pb2.GenerateCaption:
        caption = hskw_pb2.GenerateCaption()
        caption.ParseFromString(message.body)

        return caption

    async def process_message(
        self, message: aio_pika.abc.AbstractIncomingMessage
    ) -> None:
        async with message.process():
            start = default_timer()
            request = self.parse_message(message)
            log.info("Received caption request", caption=request)

            image_buf = self.storage.get_image(request.image_path)
            generated_caption = self.model.inference(image_buf)

            async with grpc.aio.insecure_channel(self.grpc_endpoint) as channel:
                stub = hskw_pb2_grpc.SnappyStub(channel)
                await stub.CreatedCaption(
                    hskw_pb2.CreatedCaptionRequest(
                        challenge_id=request.challenge_id, caption=generated_caption
                    )
                )
                log.info(
                    "Sent caption to server",
                    caption=generated_caption,
                    time_in_s=default_timer() - start,
                )


class QNAQueue(AbstractQueue):
    def __init__(self, config: QueueConfig, qa_model: QuestionAndAnswer, grpc_endpoint: str):
        super().__init__(config)
        self.model = qa_model
        self.grpc_endpoint = grpc_endpoint

    def parse_message(
        self, message: aio_pika.abc.AbstractIncomingMessage
    ) -> hskw_pb2.GenerateQuestionAnswer:
        qa = hskw_pb2.GenerateQuestionAnswer()
        qa.ParseFromString(message.body)

        return qa

    async def process_message(
        self, message: aio_pika.abc.AbstractIncomingMessage
    ) -> None:
        async with message.process():
            start = default_timer()
            request = self.parse_message(message)
            log.info("Received caption request", qa=request)

            generated_qas = self.model.inference(request.caption)
            qas = [
                hskw_pb2.QNA(question=qa.question, answer=qa.answer)
                for qa in generated_qas
            ]

            async with grpc.aio.insecure_channel(self.grpc_endpoint) as channel:
                stub = hskw_pb2_grpc.SnappyStub(channel)
                await stub.CreatedQuestion(
                    hskw_pb2.CreatedQuestionRequest(
                        challenge_id=request.challenge_id, qnas=qas
                    )
                )
                log.info(
                    "Sent generated questions to server",
                    qas=qas,
                    time_in_s=default_timer() - start,
                )
