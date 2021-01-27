package api

import "github.com/marcoshuck/raft/pkg/core"

// LogProcessor has a method to process a log.
type LogProcessor interface {
	// Process receives a log and performs operations over core.Store's state.
	Process(log core.Log) error
}
