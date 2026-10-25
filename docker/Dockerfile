FROM maheshivara/sf3000-toolchain:latest

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
  apt-get install -y --no-install-recommends \
  curl \
  && rm -rf /var/lib/apt/lists/*

ARG GO_VERSION=1.27.3
RUN set -eux; \
  case "$(uname -m)" in \
  x86_64)  goarch=amd64 ;; \
  aarch64) goarch=arm64 ;; \
  *) echo "unsupported host arch: $(uname -m)" >&2; exit 1 ;; \
  esac; \
  curl -fsSL "https://go.dev/dl/go${GO_VERSION}.linux-${goarch}.tar.gz" -o /tmp/go.tgz; \
  tar -C /usr/local -xzf /tmp/go.tgz; \
  rm /tmp/go.tgz; \
  /usr/local/go/bin/go version

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOTOOLCHAIN=local
ENV GOPATH=/tmp/go
ENV GOCACHE=/tmp/go-cache
ENV GOMODCACHE=/tmp/go-mod

COPY ./docker/mips-gcc /usr/local/bin/mips-gcc
RUN sed -i 's/\r$//' /usr/local/bin/mips-gcc && chmod 755 /usr/local/bin/mips-gcc

WORKDIR /2048

COPY ./go.* /2048/
RUN go mod download

COPY ./internal /2048/internal
COPY ./main.go /2048/main.go
COPY ./port /2048/port
