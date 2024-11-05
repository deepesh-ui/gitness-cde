# Use a base image with minimal dependencies
FROM ubuntu:22.04

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive \
    PATH=$PATH:/usr/local/go/bin:/go/bin \
    GOROOT=/usr/local/go \
    GOPATH=/go

# Install dependencies
RUN apt-get update && apt-get install -y \
    curl \
    build-essential \
    git \
    wget \
    unzip \
    openssl \
    vim \
    gnupg \
    software-properties-common

# Install Homebrew (Linuxbrew) dependencies and Homebrew itself
RUN apt-get install -y build-essential procps file && \
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)" && \
    echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"' >> /root/.bashrc && \
    eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"

# Install specific version of protobuf (v3.21.11)
RUN eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" && \
    brew install wget && \
    wget -q https://github.com/protocolbuffers/protobuf/releases/download/v3.21.11/protobuf-all-3.21.11.zip && \
    unzip protobuf-all-3.21.11.zip && \
    cd protobuf-3.21.11 && \
    ./configure && make && make install && \
    rm -rf ../protobuf-all-3.21.11.zip

# Install Go 1.20
RUN wget -q https://go.dev/dl/go1.20.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.0.linux-amd64.tar.gz && \
    rm go1.20.0.linux-amd64.tar.gz

# Install Go tools: protoc-gen-go and protoc-gen-go-grpc
RUN /usr/local/go/bin/go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
    /usr/local/go/bin/go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

# Install Node.js and Yarn
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs && \
    npm install -g yarn

# Set Go paths
ENV PATH="${PATH}:${GOPATH}/bin"

# Set the working directory
WORKDIR /workspace

# Expose port for local server
EXPOSE 3000

# Command to ensure Docker image has necessary dependencies
CMD [ "bash" ]
