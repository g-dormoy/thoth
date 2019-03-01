package server

import (
	"bufio"
	"errors"
	"log"
	"net"
)

// ConnectionHandler are functions used to handle connections queries
type ConnectionHandler func(net.Conn, ICommandDispatcher)

// ListenerFactory defines a factory to build new net.Listener
// It could be nothing more then net.Listen()
type ListenerFactory func(string, string) (net.Listener, error)

// HandleConn offers a defauilt connection handler for the thoth server
func HandleConn(c net.Conn, dispatcher ICommandDispatcher) {
	sc := bufio.NewScanner(c)

	for sc.Scan() {
		resp := Response{}
		req, err := NewRequest(sc.Text())
		if err != nil {
			resp.ErrT = ERROR_NO_COMMAND_FOUND
			resp.Err = errors.New("no command found")
		} else {
			resp = dispatcher.Dispatch(*req)
		}
		c.Write([]byte(resp.Payload()))
	}
}

// Starts a new Thoth server with its own ConnectionHandler
func Start(f ListenerFactory, ch ConnectionHandler, dispatcher ICommandDispatcher, ntwk, path string) error {
	ln, err := f(ntwk, path)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to Accept")
			continue
		}

		go ch(conn, dispatcher)
	}
}
