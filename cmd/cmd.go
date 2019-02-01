package cmd

import "fmt"

// Commander interface describes the interface for making a cmd on thoth project
type Commander interface {
	Run(string) bool
	Err() error
}

// NewCommand is a factory for the command
func NewCommand(c string) (Commander, error) {
	switch c {
	case "produce":
		return &ProduceCmd{}, nil
	}

	return nil, fmt.Errorf("Command %s not found", c)
}
