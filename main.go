package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	router := gin.New()

	// Add middleware to recover from panics
	router.Use(gin.Recovery())

	// Define the root endpoint that returns the ASCII animation
	router.GET("/", func(c *gin.Context) {
		// Set headers for streaming response
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")

		// Stream the response
		c.Stream(func(w io.Writer) bool {
			// Infinite loop for continuous animation
			for {
				// Loop through frames 1-20
				for i := 1; i <= 20; i++ {
					// Read the frame file
					frameContent, err := readFrame(i)
					if err != nil {
						fmt.Printf("Error reading frame %d: %v\n", i, err)
						continue
					}

					// Clear screen ANSI escape sequence
					clearScreen := "\033[2J\033[H"

					// Send the frame with clear screen command
					if _, err := w.Write([]byte(clearScreen + frameContent)); err != nil {
						fmt.Printf("Error writing frame %d: %v\n", i, err)
						return false
					}

					// Flush the response
					if flusher, ok := w.(http.Flusher); ok {
						flusher.Flush()
					}
					// Wait before sending next frame (reduced to 0.1 seconds)
					const frameDelay = 100 * time.Millisecond
					time.Sleep(frameDelay)
				}
			}
			// This return statement will never be reached due to the infinite loop
			// but is required by the c.Stream function signature
			return false //nolint:govet
		})
	})

	// Start server on port 8080
	fmt.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func readFrame(frameNumber int) (string, error) {
	filename := fmt.Sprintf("frames/%02d.txt", frameNumber)
	content, err := readFile(filename)
	if err != nil {
		return "", err
	}
	return content, nil
}

func readFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
