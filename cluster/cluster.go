package cluster

// Clusterable interface define the API if a thoth Cluster
type Clusterable interface {
	AddNode(Nodable)
}

// Cluster is a Clusterable struct
type Cluster struct {
	Nodes  NodeIterator
	Server Servable
}

// AddNode add a node to the cluster
func (c *Cluster) AddNode(n Nodable) {
	if c.Nodes == nil {
		c.Nodes = make(Nodes, 0)
	}

	for node := range c.Nodes.Iter() {
		if node.String() == n.String() && node.Network() == n.Network() {
			return
		}
	}
	c.Nodes = c.Nodes.Append(n.Node())
}

// New return a new instance of a cluster stuct
func New(s Servable, n NodeIterator) Cluster {
	return Cluster{
		Nodes:  n,
		Server: s,
	}
}
