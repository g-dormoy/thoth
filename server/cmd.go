package server

import (
	"fmt"
)

type CommandHandler func(IRequest) Response

// ICommandDispatcher defines and interface to add and manage command in a thot server
type ICommandDispatcher interface {
	Add(string, CommandHandler) error
	Dispatch(IRequest) Response
}

// CommandDispatcher implements ICommandDispatcher
type CommandDispatcher struct {
	cmd map[string]CommandHandler
}

// NewCommandDispatcher is CommandDispatcher factory
func NewCommandDispatcher() CommandDispatcher {
	return CommandDispatcher{
		cmd: make(map[string]CommandHandler),
	}
}

// Add a new cmd handler to the Dispatcher
// Return an error if command already exist
func (cd CommandDispatcher) Add(cmd string, handler CommandHandler) error {
	if _, ok := cd.cmd[cmd]; ok {
		return fmt.Errorf("command %s already exist", cmd)
	}
	cd.cmd[cmd] = handler
	return nil
}

// Dispatch parse the Request and send it to the write CommandHandler
func (cd CommandDispatcher) Dispatch(req IRequest) Response {
	if _, ok := cd.cmd[req.Command()]; false == ok {
		return Response{
			Err:  fmt.Errorf("%s command not found", req.Command()),
			ErrT: ERROR_STATUS_COMMAND_NOT_FOUND,
		}
	}
	return cd.cmd[req.Command()](req)
}

// SayHi is a simple command to say hi :D
func SayHi(cmd IRequest) Response {
	return Response{
		Pl: "Ho hello there!",
	}
}
