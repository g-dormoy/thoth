package server

// This file contains the definitions of the server cmd

// Event describe a payload
type Event struct {
	Label   string `json:"label"`
	TTL     uint   `json:"ttl"`
	Payload string `json:"payload"`
}

// Command struct defines a command on Thoth
type Command struct {
	cmd   string
	event Event
}
