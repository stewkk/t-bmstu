#!/usr/bin/env sh

git -C t-bmstu/ pull \
    && docker compose -f t-bmstu/deploy/docker-compose.yaml pull t-bmstu \
    && docker compose -f t-bmstu/deploy/docker-compose.yaml up --no-deps -d t-bmstu
