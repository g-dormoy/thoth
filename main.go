package main

import "sync"

func main() {
	sc := &ServiceCollection{}
	m := &sync.Mutex{}

	s := NewServer(m, sc)
	s.Routes()
	s.Start("4242")
}
