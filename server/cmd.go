package server

import (
	"fmt"
	"strings"
)

type CommandHandler func(ThothRequest) ThothResponse

// SayHi is a simple command to say hi :D
func SayHi(cmd string) ThothResponse {
	return "Ho hello there!"
}

// DispatchCmd parse the thotRequest and send it to the write CommandHandler
// DispatchCmd should be a CommandHandler itself
func DispatchCmd(query string) ThothResponse {
	var resp ThothResponse
	command := strings.SplitN(query, " ", 1)
	if command == nil {
		return "no command found"
	}
	switch cmd := command[0]; cmd {
	case "hi":
		resp = SayHi("")
	default:
		resp = ThothResponse(fmt.Sprintf("command '%s' not found", cmd))
	}
	return resp
}
