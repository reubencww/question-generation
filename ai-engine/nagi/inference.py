import io
from dataclasses import dataclass
from typing import Optional

import torch
from blip.models.blip import blip_decoder
from blip.models.blip_vqa import BLIP_VQA, blip_vqa
from PIL import Image
from question_generation.pipelines import pipeline
from torchvision import transforms
from torchvision.transforms.functional import InterpolationMode

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")


class Caption:
    """
    Image captioning with BLIP.
    """

    def __init__(
            self,
            image_size: int = 480,
            checkpoint: str = "checkpoints/model_base_capfilt_large.pth",
    ):
        self.image_size = image_size
        self.model_caption = blip_decoder(
            pretrained=checkpoint,
            image_size=image_size,
            vit="base",
        )
        self.model_caption.eval()
        self.model_caption = self.model_caption.to(device)

    def inference(self, image: bytes) -> str:
        with torch.no_grad():
            caption = self.model_caption.generate(
                load_image(image),
                sample=False,
                num_beams=3,
                max_length=20,
                min_length=5,
            )

            return caption[0]


@dataclass
class QNAResult:
    question: str
    answer: str


class QuestionAndAnswer:
    """
    Question and answer generation using the question_generation repository.
    """

    def __init__(self, ):
        self.nlp = pipeline("multitask-qa-qg")

    def inference(self, context: str) -> list[QNAResult]:
        raw_results = self.nlp(context)
        return [
            QNAResult(question=result["question"], answer=result["answer"])
            for result in [
                # deduplicate results
                i
                for n, i in enumerate(raw_results)
                if i not in raw_results[n + 1:]
            ]
        ]


class VQAModelLoader:
    """
    Preloads the VQA model. Unlike the other models, the VQA process usually involves asking multiple questions for a
    single image. The model only needs to be loaded once.

    In other words, only one VQAModelLoader should be instantiated and passed into new instances of VQA. The other
    models are "stateless", while the VQA model is "per-image".
    """

    def __init__(self, image_size: int = 480, checkpoint: str = "checkpoints/model_base_vqa_large.pth"):
        self.model: Optional[BLIP_VQA] = None
        self.image_size = image_size
        self.checkpoint = checkpoint

    def load(self):
        self.model = blip_vqa(
            pretrained=self.checkpoint,
            image_size=self.image_size,
            vit="base",
        )
        self.model.eval()
        self.model = self.model.to(device)

    def get_loaded_model(self) -> BLIP_VQA:
        if self.model is None:
            self.load()

        return self.model


class VQA:
    """
    Facilitates the VQA process with a given image. The model should be preloaded (see VQAModelLoader). The image is
    loaded only once. Powered by BLIP.
    """

    def __init__(self, model_loader: VQAModelLoader, image: bytes):
        self.model_question = model_loader.get_loaded_model()
        self.image = load_image(image)

    def ask(self, question: str) -> str:
        with torch.no_grad():
            return self.model_question(
                self.image, question, train=False, inference="generate"
            )


def load_image(img: bytes, image_size: int = 480):
    raw_image = Image.open(io.BytesIO(img)).convert("RGB")

    transform = transforms.Compose(
        [
            transforms.Resize(
                (image_size, image_size), interpolation=InterpolationMode.BICUBIC
            ),
            transforms.ToTensor(),
            transforms.Normalize(
                (0.48145466, 0.4578275, 0.40821073),
                (0.26862954, 0.26130258, 0.27577711),
            ),
        ]
    )
    image = transform(raw_image).unsqueeze(0).to(device)

    return image
