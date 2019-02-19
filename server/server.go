package server

import (
	"errors"
	"net"
)

// Serv is a struct defining a thoth server
type Serv struct {
	// Port is the port on which the server should listen
	Port string
	// Netwk is the type on connection the server use
	// It can be anything based on network handle by the listener
	Netwk string
	// listenerFactory is a getter that returns a net.Listener
	listenerFactory func(string, string) (net.Listener, error)
}

// New return a thoth server instance
// If no listener factory is given it returns an error
func New(netwk, port string, factory func(string, string) (net.Listener, error)) (*Serv, error) {
	if factory == nil {
		return nil, errors.New("listenerFactory can not be nil")
	}
	return &Serv{
		Port:            port,
		Netwk:           netwk,
		listenerFactory: net.Listen,
	}, nil
}
