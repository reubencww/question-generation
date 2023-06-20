FROM condaforge/mambaforge:latest
RUN apt update && apt install -y curl

COPY environment.yml /app/
WORKDIR /app

RUN --mount=type=cache,target=/opt/conda/pkgs mamba env create -f environment.yml
RUN conda run -n donny python -m nltk.downloader punkt
RUN conda clean -afy

COPY . /app
#RUN curl --request GET -sL \
#    --url 'https://storage.googleapis.com/sfr-vision-language-research/BLIP/models/model_base_capfilt_large.pth' \
#    --output './checkpoints/model_base_capfilt_large.pth'
#RUN curl --request GET -sL \
#    --url 'https://storage.googleapis.com/sfr-vision-language-research/BLIP/models/model_base_vqa_capfilt_large.pth' \
#    --output './checkpoints/model_base_vqa_capfilt_large.pth'

CMD ["conda", "run", "--no-capture-output", "-n", "donny", "python", "main_both.py"]
