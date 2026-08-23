#!/bin/sh
set -eu
platform=${1:-linux/arm64}
tag=${2:-shotvault:local}
docker buildx build --platform "$platform" -f benzhi.Dockerfile -t "$tag" --load .
