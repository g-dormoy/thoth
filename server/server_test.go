package server_test

import (
	"net"
	"testing"

	"github.com/g-dormoy/thoth/server"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		inPort   string
		inNetwk  string
		inListen func(string, string) (net.Listener, error)
		out      server.Serv
		outErr   bool // true if an error must be thrown

	}{
		"port: 4242": {
			inPort:   "4242",
			inNetwk:  "tcp",
			inListen: net.Listen,
			out: server.Serv{
				Port:  "4242",
				Netwk: "tcp",
			},
			outErr: false,
		},
		"port: 80": {
			inPort:   "80",
			inNetwk:  "udp",
			inListen: nil,
			out: server.Serv{
				Port:  "80",
				Netwk: "udp",
			},
			outErr: true,
		},
	}

	assert := assert.New(t)
	for _, test := range tests {
		s, err := server.New(test.inNetwk, test.inPort, test.inListen)
		if s != nil {
			assert.Equal(test.out.Port, s.Port, "They should be Equal")
			assert.Equal(test.out.Netwk, s.Netwk, "They should be Equal")
		}
		if true == test.outErr {
			assert.NotNil(err)
		}
	}
}

//func TestStart(t *testing.T) {
//
//}
