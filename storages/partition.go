package storages

// Partitioner is the interface defining Partion API
type Partitioner interface{}

// Partition is a partitioner
type Partition struct {
	topics []Topic
}

// Topic defines a topic
type Topic struct {
	identifer string
}

// NewPartition creates a new instance of a Partition
func NewPartition() *Partition {
	return &Partition{}
}

// AddTopic add a new topic to a partition
func (p *Partition) AddTopic() error {
	return nil
}
