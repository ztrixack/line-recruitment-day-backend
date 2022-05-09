package ports

import (
	"election-service/internal/core/models"
)

type ElectionRepository interface {
	FindById(int) (*models.Election, error)
	UpdateById(int, models.Json) (int, error)
}

type CandidateRepository interface {
	Find() ([]models.Candidate, error)
}

type CandidateVoteRepository interface {
	Create(models.CandidateVote) (*models.CandidateVote, error)
	Find() ([]models.CandidateVote, error)
	FindById(int) (*models.CandidateVote, error)
	IncreaseById(int) (int, error)
}

type VoterRepository interface {
	Create(models.Voter) (*models.Voter, error)
	FindByNationId(string) (*models.Voter, error)
}
