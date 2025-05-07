FROM registry.gitlab.com/arquivei/dockerimages/go:${PROJECT_GO_VERSION}-alpine AS builder

ARG BUILD_VERSION
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=auto

WORKDIR /build

## Install dependencies
COPY go.mod go.sum ./
RUN go mod download

## Build code
COPY . .
RUN echo "America/Sao_Paulo" > /etc/timezone

## Build all binaries in cmd/
RUN mkdir bin
RUN for dir in ./cmd/*; do \
      name=$(basename "$dir"); \
      go build -tags musl -ldflags="-s -w -X main.version=${BUILD_VERSION}" -o "./bin/$name" "$dir"; \
      upx -1 -qqq "$name"; \
    done

FROM registry.gitlab.com/arquivei/dockerimages/go:scratch

ENV TZ="America/Sao_Paulo"
COPY --from=builder /etc/timezone /etc/timezone
COPY --from=builder /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime

## Copy all binaries from builder
COPY --from=builder /build/bin/* ./

ENTRYPOINT [ "${API_NAME}" ]