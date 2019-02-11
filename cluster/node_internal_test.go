package cluster

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	a := Addr{
		netwk: "network",
		addr:  "address",
	}

	e := Node{
		Addr: a,
	}

	n := NewNode(a)

	assert.Equal(t, e, n, "they should be equal")
}

func TestConnect(t *testing.T) {
	assert := assert.New(t)

	n := Node{
		Addr: Addr{
			netwk: "network",
			addr:  "address",
		},
	}
	// Assert non working case
	assert.NotNil(n.Connect())

	// Assert working case
	n.addr = "127.0.0.1:4242"
	n.netwk = "tcp"

	ln, err := net.Listen(n.netwk, n.addr)
	if err != nil {
		t.Errorf("could not listen on %s", n.addr)
	}
	defer ln.Close()

	assert.Nil(n.Connect())
	assert.NotNil(n.conn)

	// Assert already opened connection
	assert.Nil(n.Connect())

	n.conn.Close()
}
