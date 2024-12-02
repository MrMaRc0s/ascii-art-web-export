#!/bin/bash

# Define the image name or tag
IMAGE_NAME="dockerized"

# Get the container ID of the running container for the specific image
CONTAINER_ID=$(docker ps -q --filter "ancestor=$IMAGE_NAME")

# Check if there is a running container for the image
if [ -z "$CONTAINER_ID" ]; then
    echo "No running container found for the image $IMAGE_NAME."
else
    echo "Stopping container $CONTAINER_ID..."
    docker stop $CONTAINER_ID
    echo "Removing container $CONTAINER_ID..."
    docker rm $CONTAINER_ID
fi

# Remove the image
echo "Removing the image $IMAGE_NAME..."
docker rmi $IMAGE_NAME

# Optionally, prune unused Docker resources (not necessary for your use case)
# docker system prune -f

echo "The container and image have been removed."
