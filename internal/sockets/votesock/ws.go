package votesock

import (
	"election-service/internal/core/models"
	"election-service/pkg/ws"
	"encoding/json"

	"github.com/antoniodipinto/ikisocket"
)

type socketWs struct{}

func NewWs() socketWs {
	return socketWs{}
}

// Retrieving objects with primary key
func (b socketWs) VoteUpdated(data models.CandidateVoteResponse) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ikisocket.Fire(ws.VoteEvent, bytes)
	return nil
}
