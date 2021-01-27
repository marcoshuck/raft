package client

import "github.com/marcoshuck/raft/pkg/pb"

type VoteRequester interface {
	pb.VoteRequesterClient
}
