package electionsock

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
func (b brokerMock) StatusUpdated(data models.ElectionResponse) error {
	args := b.Called(data)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}
