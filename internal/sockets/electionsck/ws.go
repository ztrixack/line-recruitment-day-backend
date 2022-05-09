package electionsck

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
func (b brokerWs) StatusUpdated(data models.ElectionResponse) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ikisocket.Fire(ws.ElectionEvent, bytes)
	return nil
}
