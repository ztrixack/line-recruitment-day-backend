package candidatesrv

import (
	"election-service/internal/core/models"
	"election-service/internal/core/ports"
	"election-service/internal/utils"
	"election-service/internal/utils/resp"
	"election-service/pkg"
)

type Service struct {
	repository ports.CandidateRepository
}

func New(accountRepository ports.CandidateRepository) Service {
	return Service{repository: accountRepository}
}

func (s Service) GetCandidate() models.Response {
	entities, err := s.repository.Find()
	if err != nil {
		pkg.Error(err, "get all account")
		return resp.InternalServerError
	}

	results := []models.CandidateResponse{}
	if err := utils.JsonFilter(entities, &results); err != nil {
		pkg.Error(err, "convert account: %+v", entities)
		return resp.InternalServerError
	}

	return resp.OK(results)

}
