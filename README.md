## About
Web application for teachers to upload pictures, generate questions and answers using AI libraries. These are used for oral practices between teachers and students.

## Stack
Golang, Embedded VueJs, RabbitMQ for AI job processing, AI worker containers, postgresql

## Deployment/Running

> **Warning**
>
> If you're making use of the WSL2 backend, be sure to limit the max memory in `.wslconfig`. The build process is extremely memory hungry. If you're facing random errors, WSL2 is probably running out of memory. Ensure that you've allocated enough RAM to your WSL2 instance.

You'll only need to run the `curl` commands once. These models are baked into the container image for fast(er) cold starts.
```shell
curl --request GET \
    --url 'https://storage.googleapis.com/sfr-vision-language-research/BLIP/models/model_base_capfilt_large.pth' \
    --output './ai-engine/checkpoints/model_base_capfilt_large.pth'
curl --request GET \
    --url 'https://storage.googleapis.com/sfr-vision-language-research/BLIP/models/model_base_vqa_capfilt_large.pth' \
    --output './ai-engine/checkpoints/model_base_vqa_capfilt_large.pth'

# add -d when required. Startup will take a while, so you'll probably want to observe the logs to ensure everything goes smoothly.
# use scale to define the number of ai workers you want to run
docker compose -f docker-compose.yml up

# Populate the DB schema
docker exec -it control-plane-prod "/hayate" "migrate"
```

### Multi-instances

The AI engine's containers are specifically designed to scale out without any additional configuration. Just add more workers!

```shell
docker compose -f docker-compose.yml up --scale ai-engine-caption=2 --scale ai-engine-qna=2
```

### The many shades of docker-compose
- docker-compose: For "production". Spins up 2 separate containers for QNA and captioning.
- docker-compose-single: For "production". Spins up a single container that runs QNA and captioning in 2 async tasks.
- docker-compose-without-ai: For "production". Starts up all services except the AI containers.
- docker-compose-dev: For development. Starts up minio, RabbitMQ and Postgres only.

### `linux/arm64` quirks

On aarch64/Apple Silicon, due to some `compose` quirks (at the time of writing) you may need to run the following command to force a manual build for linux/amd64. One of the dependencies has stopped offering automated builds for linux/arm only. This is not an issue when running the AI engine on macOS natively.

```shell
docker compose -f docker-compose-single.yml build ai-engine
docker compose -f docker-compose-single.yml up
```

### Native AI engine

```shell
docker compose -f docker-compose-without-ai.yml up
```

### Caveats

- RabbitMQ takes a while to start up, the apps (ai-engine/control-plane) may fail to start. They will automatically retry.
- Public buckets: in a real scenario, you'll want to use an actual service and/or not run minio on containers or on the same stack. Private buckets + presigned URLs don't play nice with minio in containers.


## Development
This starts up the following services:
- RabbitMQ
- PostgreSQL
- minio

```shell
docker compose -f docker-compose-dev.yml up
```

### AI engine

> **Note**
> If any of the commands below fail, replace `mamba` with `conda`.

- initial creation: `mamba env create -f environment.yml`
- activate the env: `conda activate donny`
- update the env: `mamba env update -f environment.yml --prune`


### Generate gRPC bindings
You'll need to have protoc installed. Run the following in the root directory.

```shell
protoc --go_out=control-plane --go_opt=paths=source_relative \
       --go-grpc_out=control-plane --go-grpc_opt=paths=source_relative \
       hskw/hskw.proto

python -m grpc_tools.protoc -I. --python_out=ai-engine --pyi_out=ai-engine --grpc_python_out=ai-engine hskw/hskw.proto
```

### AI engine Models and Checkpoints
Inside the ai-engine folder, run the following commands to download the models and checkpoints. These will be copied into the docker image when building.

```shell
conda run -n donny python -m nltk.downloader punkt
curl --request GET -sL \
    --url 'https://storage.googleapis.com/sfr-vision-language-research/BLIP/models/model_base_capfilt_large.pth' \
    --output './checkpoints/model_base_capfilt_large.pth'
curl --request GET -sL \
    --url 'https://storage.googleapis.com/sfr-vision-language-research/BLIP/models/model_base_vqa_capfilt_large.pth' \
    --output './checkpoints/model_base_vqa_capfilt_large.pth'
```
