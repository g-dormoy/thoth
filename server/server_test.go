package server_test

import (
	"errors"
	"net"
	"testing"

	"github.com/g-dormoy/thoth/server"
	"github.com/g-dormoy/thoth/testutils"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := map[string]struct {
		inPort  string
		inNetwk string
		out     server.Serv
	}{
		"port: 4242": {
			inPort:  "4242",
			inNetwk: "tcp",
			out: server.Serv{
				Port:  "4242",
				Netwk: "tcp",
			},
		},
		"port: 80": {
			inPort:  "80",
			inNetwk: "udp",
			out: server.Serv{
				Port:  "80",
				Netwk: "udp",
			},
		},
	}

	assert := assert.New(t)
	for _, test := range tests {
		s := server.New(test.inNetwk, test.inPort)
		assert.Equal(test.out.Port, s.Port, "They should be Equal")
		assert.Equal(test.out.Netwk, s.Netwk, "They should be Equal")
	}
}

func TestListen(t *testing.T) {
	tests := map[string]struct {
		listener net.Listener
		err      error
		netwk    string
		addr     string
	}{
		"When error is nil": {
			listener: testutils.NewMockListener(nil),
			err:      nil,
			netwk:    "testNetwk",
			addr:     "testAddr",
		},
		"When a error is thrown": {
			listener: nil,
			err:      errors.New("failed"),
			netwk:    "error",
			addr:     "error",
		},
	}

	assert := assert.New(t)
	for _, test := range tests {
		f := func(netwk, addr string) (net.Listener, error) {
			assert.Equal(test.netwk, netwk, "They should be Equal")
			assert.Equal(":"+test.addr, addr, "They should be Equal")
			return test.listener, test.err
		}

		s := server.Serv{
			Port:  test.addr,
			Netwk: test.netwk,
		}

		ln, err := s.Listen(f)
		if err == nil {
			assert.NotNil(ln)
		}
	}

}
