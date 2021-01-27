package api

import "github.com/marcoshuck/raft/pkg/pb"

type EntriesAppender interface {
	pb.EntriesAppenderServer
}
