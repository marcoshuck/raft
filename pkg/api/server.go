package api

import (
	"context"
	"github.com/marcoshuck/raft/pkg/core"
	"github.com/marcoshuck/raft/pkg/pb"
	"github.com/marcoshuck/raft/pkg/state"
)

// Server represents a node in RAFT. It's used to represent either a Leader and Followers.
type Server struct {
	// Status represents the status of the current RAFT node.
	Status core.Status

	// State is the
	State state.State

	pb.UnimplementedEntriesAppenderServer
	pb.UnimplementedVoteRequesterServer
}

func (s *Server) Request(ctx context.Context, request *pb.VoteRequest) (*pb.VoteResponse, error) {
	panic("implement me")
}

func (s *Server) Append(ctx context.Context, request *pb.AppendRequest) (*pb.AppendResponse, error) {
	panic("implement me")
}
