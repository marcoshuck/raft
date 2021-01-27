package core

import "testing"

func TestProcessLog(t *testing.T) {
	s := make(Store, 0)

	log := Log{
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
	}

	err := s.Process(log)
	if err != nil {
		t.Fail()
	}

	if len(s) != 3 {
		t.Fail()
	}

	v, ok := s["a"]
	if !ok || v != 1 {
		t.Fail()
	}

	v, ok = s["b"]
	if !ok || v != 2 {
		t.Fail()
	}

	v, ok = s["c"]
	if !ok || v != 3 {
		t.Fail()
	}
}
