package electionhdl

import (
	"election-service/internal/core/services/electionsrv"

	"github.com/gofiber/websocket/v2"
)

type HandlerWs interface {
	PublishStatusUpdated(*websocket.Conn) error
}

type handlerWs struct {
	service electionsrv.Service
}

func NewWs(service electionsrv.Service) handlerWs {
	return handlerWs{service: service}
}
