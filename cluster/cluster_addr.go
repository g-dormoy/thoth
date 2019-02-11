package cluster

// Addr struct represent a net.Addr implementation for the cluster pkg
type Addr struct {
	netwk string
	addr  string
}

// NewAddr return a new addr struct
func NewAddr(netwk, addr string) Addr {
	return Addr{
		netwk,
		addr,
	}
}

// String returns the addr in a string
func (c Addr) String() string {
	return c.addr
}

// Network returns the network in a string
func (c Addr) Network() string {
	return c.netwk
}
