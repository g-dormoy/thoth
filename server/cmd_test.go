package server

import (
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommandDispatcher(t *testing.T) {
	eCmd := make(map[string]CommandHandler)
	cmdDisp := NewCommandDispatcher()

	assert.Equal(t, eCmd, cmdDisp.cmd, "They should be equals")
}

func TestAdd(t *testing.T) {
	tests := map[string]struct {
		dispCmd map[string]CommandHandler
		cmd     string
		handler CommandHandler
		err     error
	}{
		"Add a new command": {
			dispCmd: make(map[string]CommandHandler),
			cmd:     "hi",
			handler: func(req IRequest) Response { return Response{Pl: "hi"} },
			err:     nil,
		},
		"Add an existing command": {
			dispCmd: map[string]CommandHandler{
				"hi": func(req IRequest) Response { return Response{Pl: "hi"} },
			},
			cmd:     "hi",
			handler: func(req IRequest) Response { return Response{Pl: "hi"} },
			err:     errors.New("command hi already exist"),
		},
	}

	assert := assert.New(t)
	for desc, test := range tests {
		log.Printf("Testing case : %s", desc)
		dispatcher := CommandDispatcher{
			cmd: test.dispCmd,
		}

		err := dispatcher.Add(test.cmd, test.handler)

		assert.Equal(test.err, err, "They should be Equal")
		assert.Equal(test.handler(Request{}).Pl, dispatcher.cmd[test.cmd](Request{}).Pl)
	}
}

func TestDispatch(t *testing.T) {
	tests := map[string]struct {
		dispatcher CommandDispatcher
		req        Request
		eResp      Response
	}{
		"Call of an existing command": {
			dispatcher: CommandDispatcher{
				cmd: map[string]CommandHandler{
					"hi": func(req IRequest) Response { return Response{Pl: "hi", Err: nil, ErrT: ""} },
				},
			},
			req:   Request{Cmd: "hi"},
			eResp: Response{Pl: "hi", Err: nil, ErrT: ""},
		},
		"Call of an no existing command": {
			dispatcher: CommandDispatcher{
				cmd: make(map[string]CommandHandler),
			},
			req:   Request{Cmd: "hi"},
			eResp: Response{Pl: "", Err: errors.New("hi command not found"), ErrT: ERROR_STATUS_COMMAND_NOT_FOUND},
		},
	}

	assert := assert.New(t)
	for desc, test := range tests {
		log.Printf("Testing case : %s", desc)
		resp := test.dispatcher.Dispatch(test.req)

		assert.Equal(test.eResp, resp, "They should be equal")
	}
}
