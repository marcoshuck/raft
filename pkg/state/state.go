package state

import (
	"encoding/json"
	"github.com/marcoshuck/raft/pkg/core"
	"os"
)

type State struct {
	Term     int64
	VotedFor int64
	Log      core.Log
	Values   map[string]float64

	CommitIndex int64
	LastApplied int64
}

func (s *State) Append(log core.Log, lastApplied, commitIndex int64) error {

	if err := s.saveValues(log); err != nil {
		return err
	}

	if err := s.updateEntries(log); err != nil {
		return err
	}

	s.increaseLastApplied(lastApplied)
	s.setCommitIndex(commitIndex)

	err := s.persistState()
	if err != nil {
		return err
	}

	return nil
}

func (s *State) saveValues(log core.Log) error {
	for _, e := range log {
		s.Values[e.Key] = e.Value
	}
	return nil
}

func (s *State) updateEntries(log core.Log) error {
	s.Log = append(s.Log, log...)
	return nil
}

func (s *State) increaseLastApplied(lastApplied int64) {
	s.LastApplied = lastApplied
}

func (s *State) setCommitIndex(commit int64) {
	s.CommitIndex = commit
}

func (s *State) persistState() error {
	file, err := os.Create("state.json")
	defer file.Close()

	if err != nil {
		return err
	}

	b, err := json.Marshal(s)
	if err != nil {
		return err
	}

	_, err = file.Write(b)
	if err != nil {
		return err
	}

	return nil
}
