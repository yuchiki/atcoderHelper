#!/bin/bash
set -eux

docker build -t ach_integration_test:latest .
docker run ach_integration_test:latest
