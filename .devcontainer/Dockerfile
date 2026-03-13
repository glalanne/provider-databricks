FROM mcr.microsoft.com/devcontainers/base:ubuntu

RUN sudo apt-get update \
	&& sudo apt-get install -y --no-install-recommends \
		git-lfs \
		golang-go \
		golang-golang-x-tools \
		jq \
		make \
		python-is-python3 \
		python3-pip \
		python3-yaml \
		ripgrep \
		shellcheck \
		tree \
		yq \
	&& sudo rm -rf /var/lib/apt/lists/*
