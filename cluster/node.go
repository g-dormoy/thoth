package cluster

import (
	"net"
)

// NodeIterator is an interface to iter through nodes
type NodeIterator interface {
	Iter() <-chan Node
	Append(Node) Nodes
	Nodes() Nodes
}

// Nodes return a copy of the Node slice
func (n Nodes) Nodes() Nodes {
	return n
}

// Iter Iterates through an Node slice
func (n Nodes) Iter() <-chan Node {
	c := make(chan Node)
	go func() {
		for _, node := range n {
			c <- node
		}
		close(c)
	}()
	return c
}

// Append a node into the Nodes slice
func (n Nodes) Append(node Node) Nodes {
	return append(n, node.Node())
}

// Nodes is a slice of Node, it implements a NodeIterator interface
type Nodes []Node

// Nodable interface defines the API for Clusterable structs
type Nodable interface {
	net.Addr
	Connect() error
	Node() Node
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

// Node returns the node
func (n Node) Node() Node {
	return n
}
