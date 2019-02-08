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
	assert.Equal(t, 1, len(c.Nodes), "they should be equal")

	c.AddNode(&n)
	assert.Equal(t, 1, len(c.Nodes), "Should not have an other Node")
}
