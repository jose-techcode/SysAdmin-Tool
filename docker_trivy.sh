#!/bin/bash

# Stop in first error

set -e

# Root use proibithed

if [ "$EUID" -eq 0 ]; then
    echo "Don't use this script with root permission."
    exit 1
fi

# Building the binary

go build -o sys

# Building the image

docker build -t sys:latest .

# Scan the docker image with trivy tool (false positives are expected)

trivy image sys:latest