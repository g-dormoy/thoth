package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func execute(cmd Command) error {
	switch cmd.cmd {
	case "PRODUCE":
		fmt.Printf("%+v", cmd.event)
	}
	return nil
}

func handleConnection(c net.Conn) error {
	defer c.Close()
	sc := bufio.NewScanner(c)

	for sc.Scan() {
		entry := strings.SplitN(sc.Text(), " ", 2)
		event := Event{}
		json.Unmarshal([]byte(entry[1]), &event)
		execute(Command{entry[0], event})
	}

	return sc.Err()
}

// Run is the entrypoint of a Thoth Server
func Run(port uint) error {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(int(port)))
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error nicely
		}
		go handleConnection(conn)
	}
}
