package candidatevoterepo

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
func (r repositoryMock) Create(entity models.CandidateVote) (*models.CandidateVote, error) {
	args := r.Called(entity)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CandidateVote), args.Error(1)
}

// Get all records
func (r repositoryMock) Find() ([]models.CandidateVote, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.CandidateVote), args.Error(1)
}

// Retrieving objects with primary key
func (r repositoryMock) FindByCandidateId(id int) (*models.CandidateVote, error) {
	args := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.CandidateVote), args.Error(1)
}

// Update with primary key
func (r repositoryMock) IncreaseByCandidateId(id int) (int, error) {
	args := r.Called(id)
	return args.Int(0), args.Error(1)
}
