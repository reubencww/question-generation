import asyncio
from os import environ

from dotenv import load_dotenv
from nagi.inference import Caption
from nagi.nagi import CaptionQueue, QueueConfig
from nagi.storage import Storage, StorageConfig


async def captioning() -> None:
    application = CaptionQueue(
        config=QueueConfig(
            address=environ.get("QUEUE_HOST", "amqp://guest:guest@localhost/"),
            queue_name=environ.get("CAPTION_QUEUE_NAME", "caption"),
        ),
        caption_model=Caption(),
        storage=Storage(
            StorageConfig(
                endpoint=environ.get("STORAGE_ENDPOINT", "localhost:9000"),
                access_key=environ.get("STORAGE_ACCESS_KEY", "minioadmin"),
                secret_key=environ.get("STORAGE_SECRET_KEY", "minioadmin"),
                bucket=environ.get("STORAGE_BUCKET_NAME", "hayate"),
            )
        ),
        grpc_endpoint=environ.get("GRPC_ENDPOINT", "localhost:8001"),
    )

    await application.run()
    try:
        await asyncio.Future()
    finally:
        await application.connection.close()


if __name__ == "__main__":
    load_dotenv()
    asyncio.run(captioning())
