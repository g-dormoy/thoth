package main

import (
	"log"
	"net"

	"github.com/g-dormoy/thoth/server"
)

func main() {
	dispatcher := server.NewCommandDispatcher()
	if err := dispatcher.Add("hi", server.SayHi); nil != err {
		log.Fatal(err)
	}
	server.Start(net.Listen, server.HandleConn, dispatcher, "tcp", ":4242")
}
