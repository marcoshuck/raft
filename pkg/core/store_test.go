package core

import (
	"github.com/matryer/is"
	"testing"
)

func TestProcessLog(t *testing.T) {
	is := is.New(t)

	// Initialize an empty store
	s := make(Store, 0)

	// Process the logs
	err := s.Process(Log{
		Entry{
			Key:   "a",
			Value: 1,
		},
		Entry{
			Key:   "b",
			Value: 2,
		},
		Entry{
			Key:   "c",
			Value: 3,
		},
	})
	is.NoErr(err)

	// Make sure it has the right length
	is.Equal(len(s), 3)

	// Check each entry has the right value
	v, ok := s["a"]
	is.Equal(v, Value(1))
	is.Equal(ok, true)

	v, ok = s["b"]
	is.Equal(v, Value(2))
	is.Equal(ok, true)

	v, ok = s["c"]
	is.Equal(v, Value(3))
	is.Equal(ok, true)

	// Run Process again
	err = s.Process(Log{
		Entry{
			Key:   "a",
			Value: 999,
		},
	})
	is.NoErr(err)

	// Check a's value again
	v, ok = s["a"]
	is.Equal(v, Value(999))
	is.Equal(ok, true)
}
