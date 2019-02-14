package cluster_test

import (
	"testing"

	"github.com/g-dormoy/thoth/cluster"
	"github.com/stretchr/testify/assert"
)

func TestAddNode(t *testing.T) {
	c := cluster.Cluster{}
	n := cluster.Node{}

	c.AddNode(&n)
	assert.Equal(t, 1, len(c.Nodes.Nodes()), "they should be equal")

	c.AddNode(&n)
	assert.Equal(t, 1, len(c.Nodes.Nodes()), "Should not have an other Node")
}

func TestNew(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		n cluster.Nodes
		s cluster.Servable
	}{
		{
			nil,
			cluster.Server{},
		}, {
			n: make(cluster.Nodes, 0),
			s: nil,
		},
	}

	for _, test := range tests {
		c := cluster.New(test.s, test.n)
		assert.Equal(c.Server, test.s, "They should be equal")
		assert.Equal(c.Nodes, test.n, "They should be equal")
	}
}
