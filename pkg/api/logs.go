package api

import "github.com/marcoshuck/raft/pkg/core"

// LogAppender has a method to append log entries.
type LogAppender interface {
	// Append receives a set of entries (log) and appends these changes over core.Store's state.
	Append(log core.Log, lastApplied, commitIndex int64) error
}
