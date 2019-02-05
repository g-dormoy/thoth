package server

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConf(t *testing.T) {
	conf := NewConf()

	assert.Equal(t, conf.port, uint(4242), "They should be equal")
	assert.Equal(t, conf.storageDir, "/var/tmp", "They should be equal")
}

func TestSetPort(t *testing.T) {
	conf := Conf{}

	ln, err := net.Listen("tcp", ":4242")
	assert.NotNil(t, conf.SetPort(uint(4242)))

	if err != nil {
		ln.Close()
	}
}
