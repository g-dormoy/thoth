package cluster

import (
	"net"
)

// Nodable interface defines the API for Clusterable structs
type Nodable interface {
	net.Addr
	Connect() error
}

// Node implement the Nodable struct, it stores a net.Conn interface
// which links the cluster to its membres
// Nodes are either alive or dead and its state is represented by the isAlive boolean
type Node struct {
	conn    net.Conn
	isAlive bool
	Addr
}

// NewNode create a Node struct
// It take a net.Addr interface as argument
// With the net.Addr interface we can make sur a Network and a address is given
func NewNode(addr net.Addr) Node {
	return Node{
		Addr: Addr{
			netwk: addr.Network(),
			addr:  addr.String(),
		},
	}
}

// Connect opens a connection to the node based on its implementation of the net.Addr interface
// If a connection is made, the conn Interface return by Dial is stored in the Node.conn field
// If a connection is already establish it does nothing
func (n *Node) Connect() error {
	if n.conn != nil {
		return nil
	}

	conn, err := net.Dial(n.Network(), n.String())
	if err != nil {
		return err
	}

	n.conn = conn
	return nil
}
