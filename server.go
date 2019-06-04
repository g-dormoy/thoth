package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Server defines the http server use to communicate with thot
type Server struct {
	router   *http.ServeMux
	mutex    *sync.Mutex
	services *ServiceCollection
}

// NewServer return a new instance of the server struct
func NewServer(m *sync.Mutex, sc *ServiceCollection) *Server {
	return &Server{
		router:   http.NewServeMux(),
		mutex:    m,
		services: sc,
	}
}

// Routes describes the handlers used by the server
func (s *Server) Routes() {
	s.router.HandleFunc("/add", addHandler(s.mutex, s.services))
	s.router.HandleFunc("/get", getHandler(s.mutex, s.services))
}

// Start an http server
func (s Server) Start(port string) {
	http.ListenAndServe(fmt.Sprintf(":%s", port), s.router)
}
