package votehdl

import (
	"election-service/internal/core/models"
	"election-service/internal/core/services/votesrv"
	"election-service/internal/utils/resp"
	"election-service/pkg"

	"github.com/gofiber/fiber/v2"
)

type HandlerRest interface {
	Vote(c *fiber.Ctx) error
	GetVoterStatus(c *fiber.Ctx) error
}

type handlerRest struct {
	voterSrv votesrv.Service
}

func NewRest(voterSrv votesrv.Service) handlerRest {
	return handlerRest{voterSrv: voterSrv}
}

func (h handlerRest) Vote(c *fiber.Ctx) error {
	var data models.CreateVoterData
	if err := c.BodyParser(&data); err != nil {
		pkg.Error(err, "convert body: %s", string(c.Body()))
		return resp.Send(c, resp.BadRequestError)
	}

	response := h.voterSrv.CreateVoter(data)

	return resp.Send(c, response)
}

func (h handlerRest) GetVoterStatus(c *fiber.Ctx) error {
	var data models.Voter
	if err := c.BodyParser(&data); err != nil {
		pkg.Error(err, "convert body: %s", string(c.Body()))
		return resp.Send(c, resp.BadRequestError)
	}

	response := h.voterSrv.GetVoterByNationId(data.NationalId)

	return resp.Send(c, response)
}
