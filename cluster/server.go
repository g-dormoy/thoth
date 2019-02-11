package cluster

import "net"

// Servable interface defines an API to build and use Thoth server
type Servable interface {
	net.Addr
}

// Server struct is a servable
type Server struct {
	Addr
}

// NewServer creates an new thoth server
func NewServer(addr net.Addr) Server {
	return Server{
		Addr{
			addr:  addr.String(),
			netwk: addr.Network(),
		},
	}
}
