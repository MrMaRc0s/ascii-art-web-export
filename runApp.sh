#!/bin/bash

# Navigate to the app directory (assuming Dockerfile is in /app)
cd ./app || { echo "Directory /app not found! Exiting."; exit 1; }

# Build the Docker image
echo "Building Docker image..."
docker build -t dockerized .

# Run the Docker container
echo "Running Docker container..."
docker run -d -p 8080:8080 --name dockerized_container dockerized

# Wait for the container to initialize (you can adjust the sleep time as needed)
echo "Waiting for the server to start..."
sleep 5

# Open the browser to localhost:8080
echo "Opening the browser..."
# For Linux/MacOS use "xdg-open" or "open" on MacOS
# For Windows, use "start" for the default browser
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    xdg-open http://localhost:8080
elif [[ "$OSTYPE" == "darwin"* ]]; then
    open http://localhost:8080
elif [[ "$OSTYPE" == "cygwin" || "$OSTYPE" == "msys" ]]; then
    start http://localhost:8080
else
    echo "Unsupported OS. Please open http://localhost:8080 manually."
fi

# Test localhost:8080
echo "Testing localhost:8080..."
curl -s http://localhost:8080 || echo "Unable to connect to localhost:8080"

# Print out instructions if needed
echo "The server is running at http://localhost:8080"
