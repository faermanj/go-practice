#!/bin/bash
set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
SRC_DIR="$(dirname "$DIR")"

version_x=$(cat version.x.txt)
version_y=$(cat version.y.txt)
version_z=$(date +%Y%m%d%H%M)
version="$version_x.$version_y.$version_z"

REGISTRY_USERNAME=${REGISTRY_USERNAME:-$USER}
IMAGE_TAG=$REGISTRY_USERNAME/go-practice:$version

cd $SRC_DIR

echo "Building $IMAGE_TAG"
docker build \
    --debug \
    --progress=plain \
    --no-cache \
    -f Containerfile  \
    -t $IMAGE_TAG \
    .

echo "Pushing $IMAGE_TAG"
docker push "$IMAGE_TAG"

echo "Image[$IMAGE_TAG] built and pushed"
