package cmd

import "encoding/json"

// ProduceCmd defines a structure
type ProduceCmd struct {
	err error
}

// ProduceEvent defines an Event
type ProduceEvent struct {
	TTL       uint   `json:"ttl"`
	Data      string `json:"data"`
	Persist   bool   `json:"persist"`
	Topic     string `json:"topic"`
	Partition string `json:"partition"`
}

// NewProduceEvent create a new event from a json with default values
func NewProduceEvent(eventJSON string) (*ProduceEvent, error) {
	e := &ProduceEvent{
		Partition: "default",
		Topic:     "default",
		Persist:   false,
	}

	err := json.Unmarshal([]byte(eventJSON), e)
	return e, err
}

// Run execute the command
func (cmd *ProduceCmd) Run(event string) bool {
	_, err := NewProduceEvent(event)
	if err != nil {
		cmd.err = err
		return false
	}

	return true
}

// Err return the last error register
func (cmd *ProduceCmd) Err() error {
	return cmd.err
}
