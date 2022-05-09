package votersrv

import (
	"election-service/internal/core/models"
	"election-service/internal/core/ports"
	"election-service/internal/utils"
	"election-service/internal/utils/resp"
	"election-service/pkg"
	"strconv"
)

type Service struct {
	voterRepo    ports.VoterRepository
	electionRepo ports.ElectionRepository
	voteRepo     ports.CandidateVoteRepository
	voteSock     ports.VoteSocket
}

func New(voterRepo ports.VoterRepository, electionRepo ports.ElectionRepository, voteRepo ports.CandidateVoteRepository, voteSock ports.VoteSocket) Service {
	return Service{voterRepo: voterRepo, electionRepo: electionRepo, voteRepo: voteRepo, voteSock: voteSock}
}

func (s Service) CreateVoter(data models.CreateVoterData) models.Response {
	entity := models.Voter{}
	if err := utils.JsonFilter(data, &entity); err != nil {
		pkg.Error(err, "convert voter: %+v", data)
		return resp.InternalServerError
	}

	voter, err := s.voterRepo.FindByNationId(entity.NationalId)

	if voter != nil {
		return resp.OK(models.Json{"status": "error", "message": "Already voted"})
	}

	election, err := s.electionRepo.FindById(1)

	if election == nil {
		return resp.InternalServerError
	}

	if election.Enable == false && election.State != "voting" {
		return resp.OK(models.Json{"status": "error", "message": "Election is closed"})
	}

	response, err := s.voterRepo.Create(entity)
	if err != nil {
		pkg.Error(err, "create voter: %+v", entity)
		return resp.InternalServerError
	}
	_ = response

	vote, err := s.voteRepo.FindByCandidateId(data.CandidateId)

	if vote == nil {
		s.voteRepo.Create(models.CandidateVote{
			CandidateId: data.CandidateId,
			VotedCount:  1,
		})
	} else {
		s.voteRepo.IncreaseByCandidateId(data.CandidateId)
	}

	candidateVote := models.CandidateVoteResponse{
		Id:         strconv.Itoa(data.CandidateId),
		VotedCount: vote.VotedCount + 1,
	}

	s.voteSock.VoteUpdated(candidateVote)

	return resp.OK(models.Json{"status": true})
}

func (s Service) GetVoterByNationId(nationId string) models.Response {
	entity, err := s.voterRepo.FindByNationId(nationId)
	if err != nil {
		return resp.InternalServerError
	}

	if entity == nil {
		// no id in database
		return resp.OK(models.Json{"status": true})
	}

	return resp.OK(models.Json{"status": false})

}
