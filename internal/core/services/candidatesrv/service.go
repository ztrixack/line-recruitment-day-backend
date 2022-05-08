package candidatesrv

import (
	"election-service/internal/core/models"
	"election-service/internal/core/ports"
	"election-service/internal/utils/resp"
	"election-service/pkg"
	"strconv"
)

type Service struct {
	candidateRepo ports.CandidateRepository
	voteRepo      ports.CandidateVoteRepository
}

func New(candidateRepo ports.CandidateRepository, voteRepo ports.CandidateVoteRepository) Service {
	return Service{candidateRepo: candidateRepo, voteRepo: voteRepo}
}

func (s Service) GetCandidate() models.Response {
	candidates, err := s.candidateRepo.Find()
	if err != nil {
		pkg.Error(err, "get all candidate")
		return resp.InternalServerError
	}

	votes, err := s.voteRepo.Find()
	if err != nil {
		pkg.Error(err, "get all vote")
		return resp.InternalServerError
	}

	results := []models.CandidateResponse{}

	for index, candidate := range candidates {
		result := models.CandidateResponse{}
		result.Id = strconv.Itoa(int(candidate.Id))
		result.Name = candidate.Name
		result.Dob = candidate.Dob
		result.BioLink = candidate.BioLink
		result.ImageLink = candidate.ImageLink
		result.Policy = candidate.Policy

		if index < len(votes) {
			result.VotedCount = votes[index].VotedCount
		} else {
			s.voteRepo.Create(models.CandidateVote{
				CandidateId: int(candidate.Id),
				VotedCount:  0,
			})
		}

		results = append(results, result)
	}

	return resp.OK(results)
}
