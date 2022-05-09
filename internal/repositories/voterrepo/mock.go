package voterrepo

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

// Create Record
func (r repositoryMock) Create(entity models.Voter) (*models.Voter, error) {
	args := r.Called(entity)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Voter), args.Error(1)
}

// Retrieving objects with primary key
func (r repositoryMock) FindByNationId(id string) (*models.Voter, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Voter), args.Error(1)
}

// Update with primary key
func (r repositoryMock) UpdateById(id int, entity models.Json) (int, error) {
	args := r.Called(id, entity)
	return args.Int(0), args.Error(1)
}
