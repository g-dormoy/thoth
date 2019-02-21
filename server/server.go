package server

import "net"

// Serv is a struct defining a thoth server
type Serv struct {
	// Port is the port on which the server should listen
	Port string
	// Netwk is the type on connection the server use
	// It can be anything based on network handle by the listener
	Netwk string
}

// New return a thoth server instance
// If no listener factory is given it returns an error
func New(netwk, port string) Serv {
	return Serv{
		Port:  port,
		Netwk: netwk,
	}
}

// Listen is a net.Listener factoryDecorator
// It return a listener based on the factory given as an arg
func (s Serv) Listen(f func(string, string) (net.Listener, error)) (net.Listener, error) {
	return f(s.Netwk, ":"+s.Port)
}
