package main

import (
	"fmt"
	"net/http"
)

// 1. Define the Component Interface
// 2. Concrete Component
// 3. Decorator

// Handler defines the basic interface for all handlers
type IHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// BasicHandler is the concrete component that handles requests
type BasicHandler struct{}

func (h *BasicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handling request...")
}

// LoggingDecorator is a decorator that adds logging functionality to the handler
type LoggingDecorator struct {
	handler IHandler
}

// NewLoggingDecorator creates a new logging decorator
func NewLoggingDecorator(h IHandler) *LoggingDecorator {
	return &LoggingDecorator{handler: h}
}

func (d *LoggingDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Log the request before delegating to the actual handler
	fmt.Println("Logging request: ", r.Method, r.URL.Path)

	// Call the original handler
	d.handler.ServeHTTP(w, r)
}

func main() {
	// Create the basic handler
	basicHandler := &BasicHandler{}

	// Wrap the basic handler with the logging decorator
	loggingHandler := NewLoggingDecorator(basicHandler)

	// Serve using the logging handler
	http.Handle("/", loggingHandler)

	// Start the HTTP server
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
