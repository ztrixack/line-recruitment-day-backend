package candidatehdl

import (
	"election-service/internal/core/services/candidatesrv"
	"election-service/internal/utils/resp"

	"github.com/gofiber/fiber/v2"
)

type HandlerRest interface {
	GetAll(c *fiber.Ctx) error
}

type handlerRest struct {
	service candidatesrv.Service
}

func NewRest(service candidatesrv.Service) handlerRest {
	return handlerRest{service: service}
}

func (h handlerRest) GetAll(c *fiber.Ctx) error {
	response := h.service.GetCandidate()

	return resp.Send(c, response)
}
