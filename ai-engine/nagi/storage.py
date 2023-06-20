import io
from dataclasses import dataclass
from typing import Optional

import urllib3
from minio import Minio


@dataclass
class StorageConfig:
    endpoint: str
    access_key: str
    secret_key: str
    bucket: str
    secure: bool = False


class Storage:
    def __init__(self, config: StorageConfig):
        self.client = Minio(
            endpoint=config.endpoint,
            access_key=config.access_key,
            secret_key=config.secret_key,
            secure=config.secure,
        )
        self.bucket = config.bucket

    def get_image(self, image_path: str) -> bytes:
        response = Optional[urllib3.response.HTTPResponse]
        try:
            response: urllib3.response.HTTPResponse = self.client.get_object(
                bucket_name=self.bucket, object_name=image_path
            )
            return response.read()
        finally:
            if response is not None:
                response.close()
                response.release_conn()
