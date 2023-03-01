#!/bin/sh

docker build -t ghcr.io/xavifr/tech-test-infra/backend .
docker push ghcr.io/xavifr/tech-test-infra/backend
