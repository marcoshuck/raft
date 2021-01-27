package core

import "github.com/marcoshuck/raft/pkg/pb"

// AppendRequest wraps the protobuf definition of pb.AppendRequest.
// AppendRequest is sent to request adding a certain amount of logs to the state machine.
// It's also used to request a heartbeat from followers.
type AppendRequest pb.AppendRequest

// AppendResponse wraps the protobuf definition of pb.AppendResponse.
// AppendResponse represents the response from the follower when appending a set of entries.
type AppendResponse pb.AppendResponse
