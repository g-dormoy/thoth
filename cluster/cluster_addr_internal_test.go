package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddr(t *testing.T) {
	e := Addr{
		netwk: "testnetwk",
		addr:  "testaddr",
	}

	a := NewAddr(e.netwk, e.addr)

	assert.Equal(t, e, a, "They should be equal")
}

func TestString(t *testing.T) {
	a := Addr{
		addr: "test",
	}

	assert.Equal(t, a.addr, a.String())
}

func TestNetwork(t *testing.T) {
	a := Addr{
		netwk: "test",
	}

	assert.Equal(t, a.netwk, a.Network())
}
