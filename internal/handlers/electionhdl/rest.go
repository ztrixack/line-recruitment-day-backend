package electionhdl

import (
	"election-service/internal/core/models"
	"election-service/internal/core/services/electionsrv"
	"election-service/internal/utils/resp"
	"election-service/pkg"

	"github.com/gofiber/fiber/v2"
)

type HandlerRest interface {
	GetResult(c *fiber.Ctx) error
	GetStatus(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}

type handlerRest struct {
	service electionsrv.Service
}

func NewRest(service electionsrv.Service) handlerRest {
	return handlerRest{service: service}
}

func (h handlerRest) GetResult(c *fiber.Ctx) error {
	response := h.service.GetElectionResult()

	return resp.Send(c, response)
}

func (h handlerRest) GetStatus(c *fiber.Ctx) error {
	response := h.service.GetElection()

	return resp.Send(c, response)
}

func (h handlerRest) Update(c *fiber.Ctx) error {
	var data models.UpdateElectionData
	if err := c.BodyParser(&data); err != nil {
		pkg.Error(err, "convert body: %s", string(c.Body()))
		return resp.Send(c, resp.BadRequestError)
	}

	if data.State == "voting" && data.Enable == false {
		return resp.Send(c, resp.BadRequestError)
	}

	if data.State == "closed" && data.Enable == true {
		return resp.Send(c, resp.BadRequestError)
	}

	if data.State == "" && data.Enable == true {
		data.State = "voting"
	}

	if data.State == "" && data.Enable == false {
		data.State = "closed"
	}

	response := h.service.UpdateElection(data)

	return resp.Send(c, response)
}
