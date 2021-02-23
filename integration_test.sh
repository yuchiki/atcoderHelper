#!/bin/bash
set -eux

imageId=$(docker build -q .)
docker run "$imageId"
