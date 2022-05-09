package votesck

import (
	"election-service/internal/core/models"
	"election-service/pkg/ws"
	"encoding/json"

	"github.com/antoniodipinto/ikisocket"
)

type brokerWs struct{}

func NewWs() brokerWs {
	return brokerWs{}
}

// Retrieving objects with primary key
func (b brokerWs) VoteUpdated(data models.CandidateVoteResponse) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ikisocket.Fire(ws.VoteEvent, bytes)
	return nil
}
