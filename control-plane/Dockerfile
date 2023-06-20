FROM node:18 as node
FROM golang:1.19 AS builder
RUN apt update && apt install -y git

COPY --from=node /usr/local/lib/node_modules /usr/local/lib/node_modules
COPY --from=node /usr/local/bin/node /usr/local/bin/node
RUN ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

WORKDIR /app/frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci

WORKDIR /app
COPY . .

WORKDIR /app/frontend
RUN npm run build

WORKDIR /app
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /hayate

FROM gcr.io/distroless/static AS final
COPY --from=builder --chown=nonroot:nonroot /hayate /hayate
COPY --from=builder --chown=nonroot:nonroot /app/config.prod.yaml /config.prod.yaml

ENTRYPOINT ["/hayate", "serve", "--address", "0.0.0.0:8000", "--grpc-address", "0.0.0.0:8001"]
