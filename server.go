package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Server defines the http server use to communicate with thot
type Server struct {
	router *http.ServeMux
	mutex  *sync.Mutex
}

// NewServer return a new instance of the server struct
func NewServer() *Server {
	return &Server{
		router: http.NewServeMux(),
		mutex:  &sync.Mutex{},
	}
}

// Routes describes the handlers used by the server
func (s *Server) Routes() {
	services := &ServiceCollection{}
	s.router.HandleFunc("/add", addHandler(s.mutex, services))
	s.router.HandleFunc("/get", getHandler(s.mutex, services))
}

func (s Server) Start(port string) {
	http.ListenAndServe(fmt.Sprintf(":%s", port), s.router)
}
