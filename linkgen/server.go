package linkgen

import (
	"context"
	"linkgen/store"
	"log"
	"net/http"
	"time"
)

// Server - HTTP server struct with all dependencies, anything you need to use inside your handlers need to be attached
// as dependency on the server struct
type Server struct {
	Port      string
	LinkStore store.LinkStore
	server    *http.Server
}

// Start - starts the HTTP API server on the specified port after adding all endpoints
func (s *Server) Start() error {
	router := NewRouter()
	router.addRoute("POST", "/linkgen", s.GenerateMinifiedLink)
	router.addRoute("GET", "/linkgen/:code", s.RedirectToOriginalURL)
	server := http.Server{
		Addr:              ":" + s.Port,
		Handler:           router.Serve(),
		ReadHeaderTimeout: 30 * time.Second,
	}

	s.server = &server

	log.Println("Starting server on port " + s.Port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Stop - sends signal to gracefully stop the server, this is useful to avoid losing data for requests that are being handled
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// New - creates a new instance for the LinkGen API HTTP Server
func New(port string, linkStore store.LinkStore) *Server {
	return &Server{
		Port:      port,
		LinkStore: linkStore,
	}
}
