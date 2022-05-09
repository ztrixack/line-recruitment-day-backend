package electionsock

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
func (b socketWs) StatusUpdated(data models.ElectionResponse) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	ikisocket.Fire(ws.ElectionEvent, bytes)
	return nil
}
