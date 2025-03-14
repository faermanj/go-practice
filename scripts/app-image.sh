#!/bin/bash

docker build \
    --debug \
    --progress=plain \
    --no-cache \
    -f Containerfile  \
    -t go-practice \
    .
