package server

import (
	"fmt"
	"strings"
)

// IRequest define an API to get request data on Thoth
type IRequest interface {
	Command() string
	Payload() string
}

// Request implements IRequest
type Request struct {
	Cmd string
	Pl  string
}

// Command returns the command of the request
func (req Request) Command() string {
	return req.Cmd
}

// Payload reuturns the Payload of the request
func (req Request) Payload() string {
	return req.Pl
}

// NewRequest is a Request factory
func NewRequest(query string) (*Request, error) {
	q := strings.SplitN(query, " ", 1)
	if q == nil {
		return nil, fmt.Errorf("no command found")
	}

	pl := ""
	if 2 == len(q) {
		pl = q[1]
	}

	return &Request{
		Cmd: q[0],
		Pl:  pl,
	}, nil
}
