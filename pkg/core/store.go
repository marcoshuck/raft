package core

// Key represents the identification key for a value inside the state machine.
type Key string

// Value represents the data type saved inside the state machine.
type Value float64

// Store represents a map of Key and Value. It's used to represent the state machine.
type Store map[Key]Value
