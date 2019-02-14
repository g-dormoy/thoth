package cluster_test

import (
	"fmt"
	"testing"

	"github.com/g-dormoy/thoth/cluster"
	"github.com/stretchr/testify/assert"
)

func TestIter(t *testing.T) {
	tests := map[string]struct {
		nodes       cluster.Nodes
		expectedLen int
	}{
		"Test with 3 nodes": {
			nodes:       make(cluster.Nodes, 3),
			expectedLen: 3,
		},
		"Test with nil": {
			nodes:       nil,
			expectedLen: 0,
		},
		"Test with 0 nodes": {
			nodes:       make(cluster.Nodes, 0),
			expectedLen: 0,
		},
	}

	fmt.Printf("Test cluster.Nodes.Iter\n")
	for desc, test := range tests {
		fmt.Printf("\t%s\n", desc)
		i := 0
		for _ = range test.nodes.Iter() {
			i++
		}
		assert.Equal(t, test.expectedLen, i, "They should be Equal")
	}
}

func TestNodes(t *testing.T) {
	tests := map[string]struct {
		nodes       cluster.Nodes
		expectedLen int
	}{
		"test with 1 nodes": {
			nodes:       make(cluster.Nodes, 1),
			expectedLen: 1,
		},
		"test with 2 nodes": {
			nodes:       make(cluster.Nodes, 2),
			expectedLen: 2,
		},
		"test with 0 nodes": {
			nodes:       make(cluster.Nodes, 0),
			expectedLen: 0,
		},
		"test with nil": {
			nodes:       nil,
			expectedLen: 0,
		},
	}
	fmt.Printf("Test cluster.Nodes.Nodes\n")
	for desc, test := range tests {
		fmt.Printf("\t%s\n", desc)
		assert.Equal(t, test.expectedLen, len(test.nodes.Nodes()), "They should be Equal")
	}
}

func TestAppend(t *testing.T) {
	tests := map[string]struct {
		nodes       cluster.Nodes
		node        cluster.Node
		expectedLen int
	}{
		"test with an empty nodes slice": {
			nodes:       make(cluster.Nodes, 0),
			node:        cluster.Node{},
			expectedLen: 1,
		},
		"test with none empty nodes slice": {
			nodes:       make(cluster.Nodes, 1),
			node:        cluster.Node{},
			expectedLen: 2,
		},
		"test with nil": {
			nodes:       nil,
			node:        cluster.Node{},
			expectedLen: 1,
		},
	}
	fmt.Printf("Test cluster.Nodes.Append\n")
	for desc, test := range tests {
		fmt.Printf("\t%s\n", desc)
		assert.Equal(t, test.expectedLen, len(test.nodes.Append(test.node)), "They should be Equal")
	}
}
