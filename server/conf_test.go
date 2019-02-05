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

	ln, _ := net.Listen("tcp", ":4242")
	assert.NotNil(t, conf.SetPort(uint(4242)))

	ln.Close()

	assert.Nil(t, conf.SetPort(uint(4242)))
}

func TestPort(t *testing.T) {
	conf := Conf{
		port: uint(4242),
	}

	assert.Equal(t, conf.port, uint(4242), "they should be Equal")
}

func TestSetStorage(t *testing.T) {
	conf := Conf{}

	assert.Nil(t, conf.SetStorageDir("/tmp"))

	assert.NotNil(t, conf.SetStorageDir("/idontexistsorry"))
}

func TestStorageDir(t *testing.T) {
	conf := Conf{
		storageDir: "test",
	}

	assert.Equal(t, conf.StorageDir(), "test", "they should be equal")
}
