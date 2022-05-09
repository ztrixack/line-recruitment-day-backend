package routes

import (
	"election-service/internal/handlers/candidatehdl"
	"election-service/internal/handlers/electionhdl"
	"election-service/internal/handlers/votehdl"

	"github.com/gofiber/fiber/v2"
)

func ElectionEndpoints(app *fiber.App, h electionhdl.HandlerRest) {
	api := app.Group("/api")
	api.Get("/election/result", h.GetResult)
	api.Get("/election/toggle", h.GetStatus)
	api.Post("/election/toggle", h.Update)
}

func CandidateEndpoints(app *fiber.App, h candidatehdl.HandlerRest) {
	api := app.Group("/api")

	api.Get("/candidates", h.GetAll)
}

func VoteEndpoints(app *fiber.App, h votehdl.HandlerRest) {
	api := app.Group("/api")

	api.Post("/vote", h.Vote)
	api.Post("/vote/status", h.GetVoterStatus)
}
