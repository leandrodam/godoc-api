#!/usr/bin/bash

DOCKER_BUILDKIT=1 COMPOSE_BAKE=true docker compose --env-file .env -f infra/docker/docker-compose.yml up --build -d
