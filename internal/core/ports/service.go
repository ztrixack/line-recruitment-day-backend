package ports

import (
	"election-service/internal/core/models"
)

type ElectionService interface {
	GetElectionResult() models.Response
	GetElection() models.Response
	UpdateElection(models.UpdateElectionData) models.Response
}

type CandidateService interface {
	GetCandidate() models.Response
	GetCandidateVoteById(int) models.Response
}

type VoterService interface {
	CreateVoter(models.CreateVoterData) models.Response
	GetVoter(int) models.Response
}
