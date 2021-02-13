#!/bin/bash -eu

imageId=$(docker build -q .)
docker run "$imageId"
