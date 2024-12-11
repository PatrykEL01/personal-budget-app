#!/bin/bash

set -e

APP_NAME="personal-budget"
DB_NAME="personal-budget-db"
NETWORK_NAME="personal-budget-network"
DOCKER_REGISTRY="patrykel1"
# Function to increment version tag
increment_tag() {
  local current_tag=$1
  IFS='.' read -r major minor patch <<< "$current_tag"
  patch=$((patch + 1))
  echo "${major}.${minor}.${patch}"
}

echo "Fetching the latest Docker image tag..."
LATEST_TAG=$(docker images --format "{{.Tag}}" "${DOCKER_REGISTRY}/${APP_NAME}" | sort -rV | head -n 1)

# Set the new tag
if [[ -z "$LATEST_TAG" || "$LATEST_TAG" == "<none>" ]]; then
  NEW_TAG="0.0.1" # Default starting tag
else
  NEW_TAG=$(increment_tag "$LATEST_TAG")
fi

echo "Latest tag: ${LATEST_TAG:-none}, New tag: $NEW_TAG"

echo "Creating Docker network..."
docker network inspect $NETWORK_NAME >/dev/null 2>&1 || \
docker network create $NETWORK_NAME

if docker ps -q --filter "name=^${DB_NAME}$" | grep -q .; then
  echo "PostgreSQL is already running. Skipping database startup."
else
  echo "Starting PostgreSQL..."
  docker run --rm -d \
    --name $DB_NAME \
    --network $NETWORK_NAME \
    -e POSTGRES_USER=personal_budget_user \
    -e POSTGRES_PASSWORD=personal_budget_password \
    -e POSTGRES_DB=personal_budget \
    -p 5432:5432 \
    postgres:15-alpine

  # Wait until the database is ready
  echo "Waiting for the database to be ready..."
  until docker exec $DB_NAME pg_isready -U personal_budget_user >/dev/null 2>&1; do
    sleep 1
  done
fi

echo "Generating .env file..."
DB_CONTAINER_IP=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $DB_NAME)
echo "DATABASE_URL=postgres://personal_budget_user:personal_budget_password@$DB_CONTAINER_IP:5432/personal_budget" > .env


echo "Building Docker image with tag: ${NEW_TAG}..."
docker build -t "${DOCKER_REGISTRY}/${APP_NAME}:${NEW_TAG}" .

# Run the Go application
echo "Starting the Go application..."
docker run --rm -d \
  --name $APP_NAME \
  --network $NETWORK_NAME \
  --env-file .env \
  -p 8080:8080 \
  "${DOCKER_REGISTRY}/${APP_NAME}:${NEW_TAG}"

echo "Environment successfully set up!"
echo "Go application is available at: http://localhost:8080"
