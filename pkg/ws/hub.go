package ws

type Hub struct {
	Election map[string]string
	Vote     map[string]string
}

var (
	ElectionEvent = "ELECTION.STATUS"
	VoteEvent     = "VOTE.COUNT"
)
