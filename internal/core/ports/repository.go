package ports

import (
	"election-service/internal/core/models"
)

type ElectionRepository interface {
	FindByID(int) (*models.Election, error)
	UpdateByID(int, models.Json) (int, error)
}

type CandidateRepository interface {
	Find() ([]models.Candidate, error)
}

type CandidateVoteRepository interface {
	FindByID(int) (*models.CandidateVote, error)
	UpdateByID(int, models.Json) (int, error)
}

type VoterRepository interface {
	Create(models.Voter) (*models.Voter, error)
	FindByID(int) (*models.Voter, error)
}
