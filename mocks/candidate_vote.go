package mocks

import (
	"election-service/internal/core/models"
)

var CandidateVote = models.CandidateVote{
	CandidateId: 8,
	VotedCount:  0,
}

var CandidateVotes = []models.CandidateVote{
	{
		CandidateId: 1,
		VotedCount:  1765,
	},
	{
		CandidateId: 2,
		VotedCount:  543,
	},
	{
		CandidateId: 3,
		VotedCount:  32,
	},
	{
		CandidateId: 4,
		VotedCount:  2212,
	},
	{
		CandidateId: 5,
		VotedCount:  33,
	},
	{
		CandidateId: 6,
		VotedCount:  4325,
	},
	{
		CandidateId: 7,
		VotedCount:  212,
	},
	{
		CandidateId: 8,
		VotedCount:  765,
	},
}
