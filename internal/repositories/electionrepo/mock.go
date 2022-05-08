package electionrepo

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

// Retrieving objects with primary key
func (r repositoryMock) FindByID(id int) (*models.Election, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Election), args.Error(1)
}

// Update with primary key
func (r repositoryMock) UpdateByID(id int, entity models.Json) (int, error) {
	args := r.Called(id, entity)
	return args.Int(0), args.Error(1)
}
