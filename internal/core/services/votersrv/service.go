package votersrv

import (
	"election-service/internal/core/models"
	"election-service/internal/core/ports"
	"election-service/internal/utils"
	"election-service/internal/utils/resp"
	"election-service/pkg"

	"github.com/gookit/validate"
)

type Service struct {
	repository ports.VoterRepository
}

func New(voterRepository ports.VoterRepository) Service {
	return Service{repository: voterRepository}
}

func (s Service) CreateVoter(data models.CreateVoterData) models.Response {
	v := validate.Struct(&data)
	if !v.Validate() {
		pkg.Error(v.Errors, "validate voter: %+v", data)
		return resp.UnprocessableEntity
	}

	entity := models.Voter{}
	if err := utils.JsonFilter(data, &entity); err != nil {
		pkg.Error(err, "convert voter: %+v", data)
		return resp.InternalServerError
	}

	response, err := s.repository.Create(entity)
	if err != nil {
		pkg.Error(err, "create voter: %+v", entity)
		return resp.InternalServerError
	}

	result := models.VoterResponse{}
	if err := utils.JsonFilter(response, &result); err != nil {
		pkg.Error(err, "convert voter: %+v", response)
		return resp.InternalServerError
	}

	return resp.Created(result)
}

func (s Service) GetVoter(id int) models.Response {
	entity, err := s.repository.FindByID(id)
	if err != nil {
		pkg.Error(err, "get one voter: %d", id)
		return resp.InternalServerError
	}

	if entity == nil {
		return resp.NotFoundError
	}

	result := models.VoterResponse{}
	if err := utils.JsonFilter(entity, &result); err != nil {
		pkg.Error(err, "convert voter: %+v", entity)
		return resp.InternalServerError
	}

	return resp.OK(result)

}
