package candidatevotesrv

import (
	"election-service/internal/core/models"
	"election-service/internal/core/ports"
	"election-service/internal/utils"
	"election-service/internal/utils/resp"
	"election-service/pkg"
)

type Service struct {
	repository ports.CandidateVoteRepository
}

func New(candidateVoteRepository ports.CandidateVoteRepository) Service {
	return Service{repository: candidateVoteRepository}
}

func (s Service) GetCandidateVote(id int) models.Response {
	entity, err := s.repository.FindByID(id)
	if err != nil {
		pkg.Error(err, "get one candidate vote: %d", id)
		return resp.InternalServerError
	}

	if entity == nil {
		return resp.NotFoundError
	}

	result := models.CandidateVoteResponse{}
	if err := utils.JsonFilter(entity, &result); err != nil {
		pkg.Error(err, "convert candidate vote: %+v", entity)
		return resp.InternalServerError
	}

	return resp.OK(result)
}

func (s Service) UpdateCandidateVote(id int, data models.UpdateCandidateVoteData) models.Response {
	entity := models.Json{}
	if err := utils.JsonFilter(data, &entity); err != nil {
		pkg.Error(err, "convert candidate vote: %+v", data)
		return resp.InternalServerError
	}

	count, err := s.repository.UpdateByID(id, entity)
	if err != nil {
		pkg.Error(err, "update candidate vote by Id: %d data: %+v", id, data)
		return resp.InternalServerError
	}

	if count != 1 {
		return resp.NotFoundError
	}

	return resp.NoContent
}
