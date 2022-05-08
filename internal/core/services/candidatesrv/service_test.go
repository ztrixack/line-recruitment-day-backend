package candidatesrv_test

import (
	"election-service/internal/core/models"
	"election-service/internal/core/services/candidatesrv"
	"election-service/internal/repositories/candidaterepo"
	"election-service/internal/repositories/candidatevoterepo"
	"election-service/internal/utils/resp"
	"election-service/mocks"
	"election-service/pkg"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService_GetCandidate(t *testing.T) {
	pkg.InitLog()
	assertions := require.New(t)

	type testcase struct {
		name     string
		expected models.Response
	}

	t.Run("GetCandidate#1", func(t *testing.T) {
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		candidateRepo.On("Find").Return(mocks.Candidates, nil)
		candidateVoteRepo.On("Find").Return(mocks.CandidateVotes, nil)
		service := candidatesrv.New(candidateRepo, candidateVoteRepo)

		expected := resp.OK(nil)

		t.Run("Success", func(t *testing.T) {
			actual := service.GetCandidate()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})

	})

	t.Run("GetCandidate#2", func(t *testing.T) {
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		candidateRepo.On("Find").Return(nil, errors.New(""))
		candidateVoteRepo.On("Find").Return(mocks.CandidateVotes, nil)
		service := candidatesrv.New(candidateRepo, candidateVoteRepo)

		expected := resp.InternalServerError

		t.Run("Candidate Repo Error", func(t *testing.T) {
			actual := service.GetCandidate()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})
	})

	t.Run("GetCandidate#3", func(t *testing.T) {
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		candidateRepo.On("Find").Return(mocks.Candidates, nil)
		candidateVoteRepo.On("Find").Return(nil, errors.New(""))
		service := candidatesrv.New(candidateRepo, candidateVoteRepo)

		expected := resp.InternalServerError

		t.Run("Vote Repo Error", func(t *testing.T) {
			actual := service.GetCandidate()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})
	})

	t.Run("GetCandidate#4", func(t *testing.T) {
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		candidateRepo.On("Find").Return(mocks.Candidates, nil)
		candidateVoteRepo.On("Find").Return(mocks.CandidateVotes[1:], nil)
		candidateVoteRepo.On("Create", mocks.CandidateVote).Return(&mocks.CandidateVote, nil)
		service := candidatesrv.New(candidateRepo, candidateVoteRepo)

		expected := resp.OK(nil)

		t.Run("Vote less than candidate", func(t *testing.T) {
			actual := service.GetCandidate()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})
	})
}
