import asyncio

import caption
import qna
from dotenv import load_dotenv


async def task_runner():
    caption_task = asyncio.create_task(caption.captioning())
    qna_task = asyncio.create_task(qna.question_answering())
    await asyncio.gather(caption_task, qna_task)


if __name__ == "__main__":
    load_dotenv()
    asyncio.run(task_runner())
