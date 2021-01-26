package core

import (
	"math/rand"
	"time"
)

// Timeout represents an election timeout used to ensure that split votes are rare and that they are resolved quickly.
type Timeout time.Duration

// NewTimeout initializes a new Timeout with value set between 150ms and 300ms as suggested on the RAFT paper.
func NewTimeout() Timeout {
	ms := rand.Int63n(150) + 150
	return Timeout(time.Duration(ms) * time.Millisecond)
}
