package candidaterepo

import (
	"election-service/internal/core/models"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func NewMock() repositoryMock {
	return repositoryMock{}
}

// Get all records
func (r repositoryMock) Find() ([]models.Candidate, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Candidate), args.Error(1)
}
