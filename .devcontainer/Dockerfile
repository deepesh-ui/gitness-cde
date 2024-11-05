# Use a base image like Ubuntu or one of Microsoft's devcontainers
FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Install Homebrew and protobuf
RUN apt-get update && apt-get install -y curl build-essential \
    && /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)" \
    && eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)" \
    && brew install protobuf \
    && apt-get clean -y && rm -rf /var/lib/apt/lists/*
