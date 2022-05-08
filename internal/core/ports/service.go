package ports

import (
	"election-service/internal/core/models"
)

type ElectionService interface {
	GetElection() models.Response
	UpdateElection(models.UpdateElectionData) models.Response
}

type CandidateService interface {
	GetCandidate() models.Response
}

type CandidateVoteService interface {
	CreateCandidateVote(models.CreateCandidateVoteData) models.Response
	GetCandidateVote(int) models.Response
	UpdateCandidateVote(int, models.UpdateCandidateVoteData) models.Response
}

type VoterService interface {
	CreateVoter(int, models.CreateVoterData) models.Response
	GetVoter(int) models.Response
}
