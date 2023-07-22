#!/bin/bash

# Get the container ID of the running container that uses the image
container_id=$(docker ps --filter "ancestor=albion_killboard" --format "{{.ID}}")

# Copy a file from the container and save it to the current folder
docker cp $container_id:/app/sql/killboard.db ./sql/killboard.db

# Build Dockerfile
docker build --no-cache -t albion_killboard .

docker run -d -v $(pwd)/sql/killboard.db:/app/sql/killboard.db albion_killboard

# Clear all 
docker stop $container_id
docker system prune -a