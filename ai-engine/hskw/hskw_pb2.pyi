from typing import ClassVar as _ClassVar
from typing import Iterable as _Iterable
from typing import Mapping as _Mapping
from typing import Optional as _Optional
from typing import Union as _Union

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf.internal import containers as _containers

DESCRIPTOR: _descriptor.FileDescriptor

class CreatedCaptionRequest(_message.Message):
    __slots__ = ["caption", "challenge_id"]
    CAPTION_FIELD_NUMBER: _ClassVar[int]
    CHALLENGE_ID_FIELD_NUMBER: _ClassVar[int]
    caption: str
    challenge_id: int
    def __init__(
        self, challenge_id: _Optional[int] = ..., caption: _Optional[str] = ...
    ) -> None: ...

class CreatedCaptionResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class CreatedQuestionRequest(_message.Message):
    __slots__ = ["challenge_id", "qnas"]
    CHALLENGE_ID_FIELD_NUMBER: _ClassVar[int]
    QNAS_FIELD_NUMBER: _ClassVar[int]
    challenge_id: int
    qnas: _containers.RepeatedCompositeFieldContainer[QNA]
    def __init__(
        self,
        challenge_id: _Optional[int] = ...,
        qnas: _Optional[_Iterable[_Union[QNA, _Mapping]]] = ...,
    ) -> None: ...

class CreatedQuestionResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GenerateCaption(_message.Message):
    __slots__ = ["challenge_id", "image_path"]
    CHALLENGE_ID_FIELD_NUMBER: _ClassVar[int]
    IMAGE_PATH_FIELD_NUMBER: _ClassVar[int]
    challenge_id: int
    image_path: str
    def __init__(
        self, challenge_id: _Optional[int] = ..., image_path: _Optional[str] = ...
    ) -> None: ...

class GenerateQuestionAnswer(_message.Message):
    __slots__ = ["caption", "challenge_id"]
    CAPTION_FIELD_NUMBER: _ClassVar[int]
    CHALLENGE_ID_FIELD_NUMBER: _ClassVar[int]
    caption: str
    challenge_id: int
    def __init__(
        self, challenge_id: _Optional[int] = ..., caption: _Optional[str] = ...
    ) -> None: ...

class QNA(_message.Message):
    __slots__ = ["answer", "question"]
    ANSWER_FIELD_NUMBER: _ClassVar[int]
    QUESTION_FIELD_NUMBER: _ClassVar[int]
    answer: str
    question: str
    def __init__(
        self, question: _Optional[str] = ..., answer: _Optional[str] = ...
    ) -> None: ...
