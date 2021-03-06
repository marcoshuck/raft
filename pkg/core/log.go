package core

// Entry represents a data structure saved in the RAFT log.
// Entry is used to perform an operation to the store.
type Entry struct {
	Key   string
	Value float64
}

// Log represents a list of changes performed to the Store.
// Log is used to recover the Store to a previous state when a failure happens.
type Log []Entry
