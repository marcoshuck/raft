package api

import "github.com/marcoshuck/raft/pkg/pb"

type VoteRequester interface {
	pb.VoteRequesterServer
}
