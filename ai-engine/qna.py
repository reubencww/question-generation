import asyncio
import os
from os import environ

from dotenv import load_dotenv
from nagi.inference import QuestionAndAnswer
from nagi.nagi import QNAQueue, QueueConfig


async def question_answering() -> None:
    application = QNAQueue(
        config=QueueConfig(
            address=environ.get("QUEUE_HOST", "amqp://guest:guest@localhost/"),
            queue_name=environ.get("CAPTION_QNA_NAME", "qna"),
        ),
        qa_model=QuestionAndAnswer(),
        grpc_endpoint=environ.get("GRPC_ENDPOINT", "localhost:8001"),
    )

    await application.run()
    try:
        await asyncio.Future()
    finally:
        await application.connection.close()


if __name__ == "__main__":
    load_dotenv()
    asyncio.run(question_answering())
