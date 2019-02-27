package server

import (
	"bufio"
	"net"
)

// ConnectionHandler are functions used to handle connections queries
type ConnectionHandler func(net.Conn)

// ListenerFactory defines a factory to build new net.Listener
// It could be nothing more then net.Listen()
type ListenerFactory func(string, string) (net.Listener, error)

type ThothResponse string
type ThothRequest string

// HandleConn offers a defauilt connection handler for the thoth server
func HandleConn(c net.Conn) {
	sc := bufio.NewScanner(c)

	for sc.Scan() {
		resp := DispatchCmd(sc.Text())
		c.Write([]byte(resp))
	}
}

// Starts a new Thoth server with its own ConnectionHandler
func Start(f ListenerFactory, ch ConnectionHandler, ntwk, path string) error {
	ln, err := f(ntwk, path)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// I BETTER DO SOMETHING HERE
		}

		go ch(conn)
	}
}
