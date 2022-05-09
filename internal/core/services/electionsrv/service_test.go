package electionsrv_test

import (
	"election-service/internal/core/models"
	"election-service/internal/core/services/electionsrv"
	"election-service/internal/repositories/candidaterepo"
	"election-service/internal/repositories/candidatevoterepo"
	"election-service/internal/repositories/electionrepo"
	"election-service/internal/utils/resp"
	"election-service/mocks"
	"election-service/pkg"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService_GetElectionResult(t *testing.T) {
	pkg.InitLog()
	assertions := require.New(t)

	type testcase struct {
		name     string
		expected models.Response
	}

	t.Run("GetElectionResult#1", func(t *testing.T) {
		electionRepo := electionrepo.NewMock()
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		electionRepo.On("FindById", 1).Return(&mocks.SolicitElection, nil)
		candidateRepo.On("Find").Return(mocks.Candidates, nil)
		candidateVoteRepo.On("Find").Return(mocks.CandidateVotes, nil)
		service := electionsrv.New(electionRepo, candidateRepo, candidateVoteRepo)
		expected := resp.OK(nil)

		t.Run("Success", func(t *testing.T) {
			actual := service.GetElectionResult()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})
	})

	t.Run("GetElectionResult#2", func(t *testing.T) {
		electionRepo := electionrepo.NewMock()
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		electionRepo.On("FindById", 1).Return(nil, errors.New(""))
		candidateRepo.On("Find").Return(nil, errors.New(""))
		candidateVoteRepo.On("Find").Return(mocks.CandidateVotes, nil)
		service := electionsrv.New(electionRepo, candidateRepo, candidateVoteRepo)
		expected := resp.InternalServerError

		t.Run("Repo Error", func(t *testing.T) {
			actual := service.GetElectionResult()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})
	})

	t.Run("GetElectionResult#3", func(t *testing.T) {
		electionRepo := electionrepo.NewMock()
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		electionRepo.On("FindById", 1).Return(nil, nil)
		candidateRepo.On("Find").Return(mocks.Candidates, nil)
		candidateVoteRepo.On("Find").Return(nil, errors.New(""))
		service := electionsrv.New(electionRepo, candidateRepo, candidateVoteRepo)
		expected := resp.InternalServerError

		t.Run("Not Found Election", func(t *testing.T) {
			actual := service.GetElectionResult()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})
	})

	t.Run("GetElectionResult#4", func(t *testing.T) {
		electionRepo := electionrepo.NewMock()
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		electionRepo.On("FindById", 1).Return(nil, nil)
		candidateRepo.On("Find").Return(nil, nil)
		candidateVoteRepo.On("Find").Return(nil, nil)
		service := electionsrv.New(electionRepo, candidateRepo, candidateVoteRepo)
		expected := resp.InternalServerError

		t.Run("Not Found Election", func(t *testing.T) {
			actual := service.GetElectionResult()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})
	})
}

func TestService_GetElection(t *testing.T) {
	pkg.InitLog()
	assertions := require.New(t)

	type testcase struct {
		name     string
		request  int
		expected models.Response
	}

	t.Run("GetElection#1", func(t *testing.T) {
		electionRepo := electionrepo.NewMock()
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		electionRepo.On("FindById", 1).Return(&mocks.SolicitElection, nil)
		service := electionsrv.New(electionRepo, candidateRepo, candidateVoteRepo)

		expected := resp.OK(nil)

		t.Run("Success", func(t *testing.T) {
			actual := service.GetElection()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})
	})
	t.Run("GetElection#2", func(t *testing.T) {
		electionRepo := electionrepo.NewMock()
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		electionRepo.On("FindById", 1).Return(nil, nil)
		service := electionsrv.New(electionRepo, candidateRepo, candidateVoteRepo)

		expected := resp.NotFoundError

		t.Run("No Id", func(t *testing.T) {
			actual := service.GetElection()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})
	})
	t.Run("GetElection#3", func(t *testing.T) {
		electionRepo := electionrepo.NewMock()
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		electionRepo.On("FindById", 1).Return(nil, errors.New(""))
		service := electionsrv.New(electionRepo, candidateRepo, candidateVoteRepo)
		expected := resp.InternalServerError

		t.Run("Repo Error", func(t *testing.T) {
			actual := service.GetElection()
			assertions.Equal(expected.Code, actual.Code)
			assertions.Equal(expected.Success, actual.Success)
		})
	})
}

func TestService_UpdateElection(t *testing.T) {
	pkg.InitLog()
	assertions := require.New(t)

	type testcase struct {
		name     string
		request  models.UpdateElectionData
		expected models.Response
	}

	t.Run("Update Election", func(t *testing.T) {
		electionRepo := electionrepo.NewMock()
		candidateRepo := candidaterepo.NewMock()
		candidateVoteRepo := candidatevoterepo.NewMock()
		electionRepo.On("UpdateById", 1, map[string]interface{}{"enable": true, "state": "voting"}).Return(1, nil)
		electionRepo.On("UpdateById", 1, map[string]interface{}{"enable": false, "state": "closed"}).Return(1, nil)
		electionRepo.On("UpdateById", 1, map[string]interface{}{"enable": false, "state": ""}).Return(0, errors.New(""))
		electionRepo.On("UpdateById", 1, map[string]interface{}{"enable": true, "state": ""}).Return(0, nil)
		candidateRepo.On("Find").Return(mocks.Candidates, nil)
		candidateVoteRepo.On("Find").Return(mocks.CandidateVotes, nil)
		service := electionsrv.New(electionRepo, candidateRepo, candidateVoteRepo)

		cases := []testcase{
			{name: "Success", request: mocks.UpdateToVoting, expected: resp.NoContent},
			{name: "Success", request: mocks.UpdateToClosed, expected: resp.NoContent},
			{name: "Repo Error", request: models.UpdateElectionData{}, expected: resp.InternalServerError},
			{name: "No ID", request: models.UpdateElectionData{Enable: true}, expected: resp.NotFoundError},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := service.UpdateElection(c.request)
				assertions.Equal(c.expected.Code, actual.Code)
				assertions.Equal(c.expected.Success, actual.Success)
			})
		}
	})
}
