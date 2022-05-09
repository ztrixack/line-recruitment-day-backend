package electionsock

import (
	"election-service/internal/core/models"

	"github.com/stretchr/testify/mock"
)

type socketMock struct {
	mock.Mock
}

func NewMock() socketMock {
	return socketMock{}
}

// Get all records
func (b socketMock) StatusUpdated(data models.ElectionResponse) error {
	args := b.Called(data)
	if args.Get(0) == nil {
		return nil
	}
	return args.Error(0)
}
