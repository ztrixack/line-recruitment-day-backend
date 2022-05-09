package votesock

import (
	"election-service/internal/core/models"

	"github.com/stretchr/testify/mock"
)

type brokerMock struct {
	mock.Mock
}

func NewMock() brokerMock {
	return brokerMock{}
}

// Get all records
func (b brokerMock) VoteUpdated(data models.CandidateVoteResponse) error {
	args := b.Called(data)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}
