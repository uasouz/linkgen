package api

import (
	"net/http"
	"os"
)

type Server struct {
	Port     string
	shutdown chan os.Signal
}

func (s *Server) Start() {
	http.ListenAndServe(":"+s.Port, nil)

	//<-s.shutdown
	// gracefull shutdown
}

func (s *Server) Stop() {
	s.shutdown <- os.Kill
}

func (s *Server) Healthz() {

}

func New(port string) *Server {
	return &Server{Port: port}
}
