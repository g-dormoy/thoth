package main

func main() {
	s := NewServer()
	s.Routes()
	s.Start("4242")
}
