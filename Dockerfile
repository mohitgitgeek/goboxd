FROM golang:1.23-bookworm

ENV DEBIAN_FRONTEND=noninteractive
ENV PATH="/usr/local/go/bin:/go/bin:${PATH}"
WORKDIR /workspace

RUN apt-get update && apt-get install -y --no-install-recommends \
    autoconf \
    bash \
    bison \
    build-essential \
    ca-certificates \
    curl \
    flex \
    gcc \
    g++ \
    git \
    iverilog \
    libnl-route-3-dev \
    libprotobuf-dev \
    libtool \
    make \
    nodejs \
    npm \
    openjdk-17-jdk \
    pkg-config \
    protobuf-compiler \
    python3 \
    && rm -rf /var/lib/apt/lists/*

COPY external/nsjail /workspace/external/nsjail
RUN make -C /workspace/external/nsjail \
    && install -m 0755 /workspace/external/nsjail/nsjail /usr/local/bin/nsjail

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@v0.6.1

COPY go.mod go.sum* ./
RUN go mod download

COPY . .
RUN go build -o /boxd-server .

EXPOSE 8080
ENTRYPOINT ["/boxd-server"]
