package core

import "github.com/marcoshuck/raft/pkg/pb"

// VoteRequest wraps the protobuf definition of pb.VoteRequest.
// VoteRequest is sent to followers to gather votes.
type VoteRequest pb.VoteRequest

// VoteResponse wraps the protobuf definition of pb.VoteResponse.
// VoteRequest represents a vote result from a follower.
type VoteResponse pb.VoteResponse
