package server

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

// Conf is a struct containing configuration data for a Thoth server
type Conf struct {
	port       uint
	storageDir string
}

// NewConf a new instance of a ServerConfig with default values
func NewConf() *Conf {
	return &Conf{
		port:       4242,
		storageDir: "/var/tmp",
	}
}

// SetPort set the port nunber
func (c *Conf) SetPort(p uint) error {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(int(p)))
	if err != nil {
		return fmt.Errorf("Could not listen to port %d", p)
	}

	ln.Close()

	c.port = p

	return nil
}

// SetStorageDir set the storage directory
func (c *Conf) SetStorageDir(path string) error {
	f, err := os.Stat(path)
	if nil != err || false == f.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	c.storageDir = path

	return nil
}

// StorageDir return the path of the storage directory
func (c *Conf) StorageDir() string {
	return c.storageDir
}
