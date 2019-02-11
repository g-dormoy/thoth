package cluster

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	a := Addr{
		addr:  "testaddr",
		netwk: "testnetwk",
	}

	e := Server{
		a,
	}

	s := NewServer(a)

	assert.Equal(t, e, s, "They should be Equal")
}
