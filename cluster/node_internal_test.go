package cluster

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	assert := assert.New(t)

	n := Node{
		Netwk: "network",
		Addr:  "address",
	}
	// Assert non working case
	assert.NotNil(n.Connect())

	// Assert working case
	n.Addr = "127.0.0.1:4242"
	n.Netwk = "tcp"

	ln, err := net.Listen(n.Netwk, n.Addr)
	if err != nil {
		t.Errorf("could not listen on %s", n.Addr)
	}
	defer ln.Close()

	assert.Nil(n.Connect())
	assert.NotNil(n.conn)

	// Assert already opened connection
	assert.Nil(n.Connect())

	n.conn.Close()
}
