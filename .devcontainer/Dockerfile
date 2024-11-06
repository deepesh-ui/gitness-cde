# Use an Ubuntu base image
FROM mcr.microsoft.com/vscode/devcontainers/base:ubuntu-22.04

# Install dependencies
RUN apt-get update && \
    apt-get install -y protobuf-compiler golang-go nodejs npm && \
    npm install -g yarn && \
    sudo rm -rf /var/lib/apt/lists/*

# Set up Go path
ENV PATH="$PATH:/go/bin"
