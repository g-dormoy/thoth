package cluster

// Clusterable interface define the API if a thoth Cluster
type Clusterable interface {
	AddNode(Nodable)
}

// Cluster is a Clusterable struct
type Cluster struct {
	Nodes []Nodable
}

// AddNode add a node to the cluster
func (c *Cluster) AddNode(n Nodable) {
	for _, node := range c.Nodes {
		if node.String() == n.String() && node.Network() == n.Network() {
			return
		}
	}
	c.Nodes = append(c.Nodes, n)
}
