package core

// Status represents the status of a certain server.
type Status string

var (
	// StatusLeader is used to mark a server as leader.
	StatusLeader Status = "leader"
	// StatusFollower is used to mark a server as follower of a certain leader.
	StatusFollower Status = "follower"
	// StatusCandidate is used to mark a server as candidate for the next election.
	StatusCandidate Status = "candidate"
	// StatusUnknown is used to represent when a server has an incorrect status.
	StatusUnknown Status = "unknown"
)
