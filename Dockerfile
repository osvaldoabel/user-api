# Base API Product
#
# VERSION               1.0.0

FROM golang:1.20.4-alpine3.17 AS base_builder
MAINTAINER Isael Sousa <faelp22@gmail.com>

WORKDIR /myapp/

COPY ["go.mod", "go.sum", "./"]

RUN go mod download


### Build Go
FROM base_builder AS builder

WORKDIR /myapp/

COPY . .

ARG PROJECT_VERSION=1 CI_COMMIT_SHORT_SHA=1
RUN go build -ldflags="-s -w -X 'main.VERSION=$PROJECT_VERSION' -X main.COMMIT=$CI_COMMIT_SHORT_SHA" -o app cmd/orch/main.go


### Build Docker Image
FROM alpine:3.17

WORKDIR /app/

COPY --from=builder /myapp/app .

ENTRYPOINT ["./app"]

#export PROJECT_VERSION=$(cat $(pwd)/VERSION)
#export CI_COMMIT_SHORT_SHA=$(git rev-parse --short HEAD) ou pegar a $CI_COMMIT_SHORT_SHA do gitlab
#docker build --build-arg PROJECT_VERSION=$(cat $(pwd)/VERSION) --build-arg CI_COMMIT_SHORT_SHA=$(git rev-parse --short HEAD) -t faelp22/product:$(cat $(pwd)/VERSION) -t faelp22/product:latest . && docker compose up -d
