package main

import (
	"net"

	"github.com/g-dormoy/thoth/server"
)

func main() {
	server.Start(net.Listen, server.HandleConn, "tcp", ":4242")
}
