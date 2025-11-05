# curl.live - ASCII Animation Server

This is a simple Go web server that serves ASCII art animations when accessed via curl.

## How it works

When you curl the server endpoint, it returns frames of ASCII art in sequence with 0.1 second delays between each frame, creating a smooth animation effect. The animation loops continuously until you interrupt it with Ctrl+C. Each frame is stored as a separate text file in the `frames/` directory.

## Prerequisites

- Go 1.25.3 or later (for local development)
- Docker (for containerized deployment)

## Running Locally

1. Clone this repository
2. Navigate to the project directory
3. Run the following commands:

```bash
# Download dependencies
go mod tidy

# Build the application
go build -o curl-live

# Run the application
./curl-live
```

The server will start on port 8080.

## Running with Docker

1. Build the Docker image:
```bash
docker build -t curl-live .
```

2. Run the container:
```bash
docker run -p 8080:8080 curl-live
```

## Testing the Animation

Once the server is running, test it with curl:

```bash
curl http://localhost:8080
```

## Project Structure

- `main.go` - The main Go application file
- `go.mod` - Go module definition
- `Dockerfile` - Docker configuration for containerization
- `.dockerignore` - Files to exclude from Docker build
- `frames/` - Directory containing the 10 ASCII art frame files

## Customizing the Animation

To create your own animation:

1. Modify the text files in the `frames/` directory
2. Each file should contain the ASCII art for one frame of your animation
3. The server will automatically read and display all files in sequence in an infinite loop

To add more frames to your animation:
1. Simply create additional text files in the `frames/` directory with sequential numbers (11.txt, 12.txt, etc.)
2. No code changes are needed - the application automatically detects all files in the directory
3. The animation will include your new frames in the sequence

## How It Works

The Go application uses the Gin web framework to serve HTTP requests. When a request is received:

1. The server sets appropriate headers for streaming text content
2. It enters an infinite loop that reads each frame file from the `frames/` directory in numerical order
3. Each frame is sent to the client with a 0.1 second delay between frames
4. ANSI escape sequences are used to clear the screen between frames for smooth animation
5. The animation continues indefinitely until the client disconnects

## License

This project is open source and available under the MIT License.
