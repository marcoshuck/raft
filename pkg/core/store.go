package core

// Key represents the identification key for a value inside the state machine.
type Key string

// Value represents the data type saved inside the state machine.
type Value float64

// Store represents a map of Key and Value. It's used to represent the state machine.
type Store map[Key]Value

// Process changes the underlying map data from changes requested in the log.
func (s Store) Process(log Log) error {
	for _, m := range log {
		s[m.Key] = m.Value
	}
	return nil
}
