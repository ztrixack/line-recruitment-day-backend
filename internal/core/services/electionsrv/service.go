package electionsrv

import (
	"election-service/internal/core/models"
	"election-service/internal/core/ports"
	"election-service/internal/utils"
	"election-service/internal/utils/resp"
	"election-service/pkg"
)

// fixed election ID
const ELECTION_ID = 1

type Service struct {
	repository ports.ElectionRepository
}

func New(ElectionRepository ports.ElectionRepository) Service {
	return Service{repository: ElectionRepository}
}

func (s Service) GetElection() models.Response {
	id := ELECTION_ID

	entity, err := s.repository.FindByID(id)
	if err != nil {
		pkg.Error(err, "get one Election: %d", id)
		return resp.InternalServerError
	}

	if entity == nil {
		return resp.NotFoundError
	}

	result := models.ElectionResponse{}
	if err := utils.JsonFilter(entity, &result); err != nil {
		pkg.Error(err, "convert Election: %+v", entity)
		return resp.InternalServerError
	}

	return resp.OK(result)

}

func (s Service) UpdateElection(data models.UpdateElectionData) models.Response {
	id := ELECTION_ID

	entity := models.Json{}
	if err := utils.JsonFilter(data, &entity); err != nil {
		pkg.Error(err, "convert Election: %+v", data)
		return resp.InternalServerError
	}

	count, err := s.repository.UpdateByID(id, entity)
	if err != nil {
		pkg.Error(err, "update Election by Id: %d data: %+v", id, data)
		return resp.InternalServerError
	}

	if count != 1 {
		return resp.NotFoundError
	}

	return resp.NoContent
}
