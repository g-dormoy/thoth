package server

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
)

func handleConnection(c net.Conn) error {
	defer c.Close()
	sc := bufio.NewScanner(c)

	for sc.Scan() {
		entry := strings.SplitN(sc.Text(), " ", 2)
		fmt.Printf("CMD NOT FOUND : %+v", entry)
	}

	return sc.Err()
}

// Bootstrap loads all the partitions into the memory
func Bootstrap(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if true == strings.Contains(file.Name(), "_partition") {
		}
	}

	return nil
}

// Run is the entrypoint of a Thoth Server
func Run(conf *Conf) error {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(int(conf.port)))
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
