package api

import (
	"linkgen/store"
	"net/http"
	"os"
)

// Server - HTTP server struct with all dependencies, anything you need to use inside your handlers need to be attached
// as dependency on the server struct
type Server struct {
	Port      string
	LinkStore store.LinkStore
	shutdown  chan os.Signal
}

// Start - starts the HTTP API server on the specified port after adding all endpoints
func (s *Server) Start() {
	http.HandleFunc("/linkgen", s.GenerateMinifiedLink)
	http.ListenAndServe(":"+s.Port, nil)

	//<-s.shutdown
	// gracefull shutdown
}

// Stop - sends signal to gracefully stop the server, this is useful to avoid losing data for requests that are being handled
func (s *Server) Stop() {
	s.shutdown <- os.Kill
}

func (s *Server) Healthz() {

}

// New - creates a new instance for the LinkGen API HTTP Server
func New(port string, linkStore store.LinkStore) *Server {
	return &Server{
		Port:      port,
		LinkStore: linkStore,
	}
}
