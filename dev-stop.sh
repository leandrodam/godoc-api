#!/usr/bin/bash

docker compose --env-file .env -f infra/docker/docker-compose.yml stop
