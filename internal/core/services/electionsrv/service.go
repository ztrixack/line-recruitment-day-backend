package electionsrv

import (
	"election-service/internal/core/models"
	"election-service/internal/core/ports"
	"election-service/internal/utils"
	"election-service/internal/utils/resp"
	"election-service/pkg"
	"fmt"
	"strconv"
)

// fixed election ID
const ELECTION_ID = 1

type Service struct {
	electionRepo  ports.ElectionRepository
	candidateRepo ports.CandidateRepository
	voteRepo      ports.CandidateVoteRepository
	electionSock  ports.ElectionSocket
}

func New(electionRepo ports.ElectionRepository, candidateRepo ports.CandidateRepository, voteRepo ports.CandidateVoteRepository, electionSock ports.ElectionSocket) Service {
	return Service{electionRepo: electionRepo, candidateRepo: candidateRepo, voteRepo: voteRepo, electionSock: electionSock}
}

func (s Service) GetElectionResult() models.Response {
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

	results := []models.ElectionResultResponse{}

	total := 0
	for _, vote := range votes {
		total += vote.VotedCount
	}

	if total == 0 {
		return resp.InternalServerError
	}

	for index, candidate := range candidates {
		result := models.ElectionResultResponse{}
		result.Id = strconv.Itoa(int(candidate.Id))
		result.Name = candidate.Name
		result.Dob = candidate.Dob
		result.BioLink = candidate.BioLink
		result.ImageLink = candidate.ImageLink
		result.Policy = candidate.Policy

		if index < len(votes) {
			result.VotedCount = votes[index].VotedCount
		}
		result.Percentage = fmt.Sprintf("%d%%", result.VotedCount*100/total)

		results = append(results, result)
	}

	return resp.OK(results)
}

func (s Service) GetElection() models.Response {
	id := ELECTION_ID

	entity, err := s.electionRepo.FindById(id)
	if err != nil {
		pkg.Error(err, "update Election by Id: %d", id)
		return resp.InternalServerError
	}

	if entity == nil {
		// return resp.NotFoundError
		return resp.InternalServerError
	}

	result := models.ElectionResponse{}
	if err := utils.JsonFilter(entity, &result); err != nil {
		pkg.Error(err, "convert voter: %+v", entity)
		return resp.InternalServerError
	}

	return resp.OK(models.Json{"status": "ok", "enable": result.Enable, "state": result.State})
}

func (s Service) UpdateElection(data models.UpdateElectionData) models.Response {
	id := ELECTION_ID

	entity := models.Json{}
	if err := utils.JsonFilter(data, &entity); err != nil {
		pkg.Error(err, "convert Election: %+v", data)
		return resp.InternalServerError
	}

	count, err := s.electionRepo.UpdateById(id, entity)
	if err != nil {
		pkg.Error(err, "update Election by Id: %d data: %+v", id, data)
		return resp.InternalServerError
	}

	if count != 1 {
		return resp.NotFoundError
	}

	response := models.ElectionResponse{
		Status: "ok",
		Enable: data.Enable,
		State:  data.State,
	}

	s.electionSock.StatusUpdated(response)

	return resp.OK(models.Json{"status": "ok", "enable": data.Enable})
}
