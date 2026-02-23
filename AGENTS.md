# curl.live Agent Guidelines

## Overview
Go web server that serves ASCII art animations via curl. Uses Gin framework for HTTP handling.

## Build & Test Commands

### Build
```bash
go build -o curl-live .
```

### Run
```bash
go run main.go
```

### Test
```bash
go test ./...
```

### Lint
```bash
golangci-lint run ./...
```

### Format
```bash
gofmt -w .
goimports -w .
```

### Dependencies
```bash
go mod tidy
```

### Docker
```bash
docker build -t curl-live .
docker run -p 8080:8080 curl-live
```

## Code Style Guidelines

### Imports
- Standard library imports first (sorted alphabetically)
- Third-party imports second (sorted alphabetically)
- Blank line between import groups
- Use explicit package names (no underscores in import paths)

### Formatting
- Use `gofmt` for all Go files
- Use `goimports` to manage imports
- 80 character line width
- Use tabs for indentation (not spaces)
- No trailing whitespace

### Naming Conventions
- Use `camelCase` for variables and functions
- Use `PascalCase` for types, interfaces, and exported symbols
- Use `UPPER_SNAKE_CASE` for constants
- Package name should be short and all lowercase
- Function names should be descriptive verbs

### Error Handling
- Always check errors from external operations (file I/O, network calls)
- Use `fmt.Printf` or logging for non-fatal errors
- Return errors to caller when appropriate
- Use `panic` only for truly exceptional situations
- The Gin framework's `Recovery()` middleware handles panics

### Types & Interfaces
- Define interfaces small and focused (max 1-3 methods)
- Prefer concrete types over interfaces for parameters
- Use `interface{}` only when necessary
- Document interface behavior with comments

### File Structure
- Single `main.go` entry point
- Helper functions in same file or separate files as needed
- Frame files in `frames/` directory with numeric names (01.txt, 02.txt, etc.)

### Error Messages
- Clear and descriptive
- Include context (e.g., "Error reading frame %d: %v")
- No emoticons or jokes in error messages

### Comments
- Use `//` for single-line comments
- Use `/* */` for multi-line comments
- Exported functions/types should have doc comments
- Keep comments up to date
- Explain "why" not "what"

### HTTP Handling (Gin)
- Use `c.Header()` for response headers
- Use `c.Stream()` for streaming responses
- Always check write errors when streaming
- Use `Flush()` to ensure data is sent immediately

### Constants
- Define constants for magic values (e.g., `const frameDelay = 100 * time.Millisecond`)
- Group related constants together

### Security
- Use non-root user in Docker (appuser)
- No hardcoded secrets or credentials
- Validate all external input
- Use `CGO_ENABLED=0` for static binaries

## Project Structure
```
curl-live/
├── main.go         # Application entry point
├── go.mod          # Go module definition
├── go.sum          # Dependency checksums
├── Dockerfile      # Container build config
├── frames/         # ASCII animation frames
│   └── 01.txt - 20.txt
└── .github/
    └── workflows/  # CI/CD configurations
```

## CI/CD
- Docker build on main branch push
- Multi-platform builds (amd64, arm64)
- Semantic versioning with auto-tagging
- Pushes to GitHub Container Registry (ghcr.io)

## Common Tasks

### Add a new frame
1. Create `frames/XX.txt` with sequential number
2. No code changes needed - frames auto-detected

### Add a new dependency
```bash
go get github.com/package/name
go mod tidy
```

### Update existing dependency
```bash
go get -u github.com/package/name
go mod tidy
```
