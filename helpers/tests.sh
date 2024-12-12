#!/bin/bash

IMAGE_NAME="personal-budget-test"
CONTAINER_NAME="personal-budget-test-container"
NETWORK_NAME="personal-budget-network"
APP_DIR=$(pwd)

echo "Building Docker image..."
docker build -t $IMAGE_NAME .

if [ $? -ne 0 ]; then
  echo "Failed to build Docker image."
  exit 1
fi

echo "Running tests in ./services in Docker container..."
docker run --rm \
  --name $CONTAINER_NAME \
  --network $NETWORK_NAME \
  --entrypoint "go" \
  -v "$APP_DIR:/app" \
  -w /app \
  $IMAGE_NAME test ./services/... -v

TEST_EXIT_CODE=$?

echo "Cleaning up..."
docker rm -f $CONTAINER_NAME 2>/dev/null

if [ $TEST_EXIT_CODE -ne 0 ]; then
  echo "Tests in ./services failed."
  exit 1
fi

echo "All tests in ./services passed successfully!"
exit 0