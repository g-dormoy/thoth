package cluster_test

import (
	"testing"

	"github.com/g-dormoy/thoth/cluster"
	"github.com/stretchr/testify/assert"
)

type addrMock struct {
	network      string
	NetworkCalls int
	address      string
	StringCalls  int
}

func (n *addrMock) Network() string {
	n.NetworkCalls++
	return n.network
}

func (n *addrMock) String() string {
	n.StringCalls++
	return n.address
}

func TestNewNode(t *testing.T) {
	na := &addrMock{
		network: "network",
		address: "address",
	}
	n := cluster.NewNode(na)

	assert.Equal(t, n.Addr, na.address, "they should be equal")
	assert.Equal(t, n.Netwk, na.network, "they should be equal")
}
