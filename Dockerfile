FROM golang:1.24-alpine AS gen_go_grpc

ARG GEN_GO_VER=1.34.2
ARG GEN_GO_GRPC_VER=1.5.1
ARG GEN_CONNECT_GO_VER=1.18.1

RUN GOBIN=/ go install google.golang.org/protobuf/cmd/protoc-gen-go@v${GEN_GO_VER}
RUN GOBIN=/ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${GEN_GO_GRPC_VER}
RUN GOBIN=/ go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v${GEN_CONNECT_GO_VER}

FROM debian:12-slim

LABEL org.opencontainers.image.source = "https://github.com/taomics/go-pramanapi"

ARG TARGETPLATFORM
ARG PROTOC_VER=28.3
ARG PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download

RUN case "${TARGETPLATFORM}" in \
      linux/amd64) \
        echo "PROTOC_ARCH=linux-x86_64" > /tmp/env; \
        ;; \
      *) exit 99 ;; \
    esac

RUN echo "protoc ${PROTOC_VER}-${TARGETPLATFORM}"

RUN apt-get update -y && \
    apt-get install wget unzip -y && \
    apt-get clean -y && \
    rm -rf /var/lib/apt/lists/*

RUN . /tmp/env; wget "${PROTOC_URL}/v${PROTOC_VER}/protoc-${PROTOC_VER}-${PROTOC_ARCH}.zip" -O protoc.zip && \
    unzip "protoc.zip" -d protoc && \
    rm "protoc.zip"

COPY --from=gen_go_grpc /protoc-gen-go /protoc-gen-go-grpc /protoc-gen-connect-go /protoc/bin/
COPY docker-entrypoint.sh /entrypoint.sh

ENV PATH=$PATH:/protoc/bin/
ENTRYPOINT ["/entrypoint.sh"]
CMD ["--help"]
