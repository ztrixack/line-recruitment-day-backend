package ports

import (
	"election-service/internal/core/models"
)

type ElectionSocket interface {
	StatusUpdated(models.ElectionResponse) error
}

type VoteSocket interface {
	VoteUpdated(models.CandidateVoteResponse) error
}
