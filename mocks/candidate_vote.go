package mocks

import (
	"election-service/internal/core/models"
	"time"

	"gorm.io/gorm"
)

var CandidateVote = models.CandidateVote{
	Model: models.Model{
		Id:        0,
		CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		DeletedAt: gorm.DeletedAt{
			Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Valid: false,
		},
	},
	CandidateId: 8,
	VotedCount:  0,
}

var CandidateVotes = []models.CandidateVote{
	{
		Model: models.Model{
			Id:        1,
			CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: false,
			},
		},
		CandidateId: 1,
		VotedCount:  10,
	},
	{
		Model: models.Model{
			Id:        2,
			CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: false,
			},
		},
		CandidateId: 2,
		VotedCount:  0,
	},
	{
		Model: models.Model{
			Id:        3,
			CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: false,
			},
		},
		CandidateId: 3,
		VotedCount:  0,
	},
	{
		Model: models.Model{
			Id:        4,
			CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: false,
			},
		},
		CandidateId: 4,
		VotedCount:  0,
	},
	{
		Model: models.Model{
			Id:        5,
			CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: false,
			},
		},
		CandidateId: 5,
		VotedCount:  0,
	},
	{
		Model: models.Model{
			Id:        6,
			CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: false,
			},
		},
		CandidateId: 6,
		VotedCount:  0,
	},
	{
		Model: models.Model{
			Id:        7,
			CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: false,
			},
		},
		CandidateId: 7,
		VotedCount:  0,
	},
	{
		Model: models.Model{
			Id:        8,
			CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			DeletedAt: gorm.DeletedAt{
				Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: false,
			},
		},
		CandidateId: 8,
		VotedCount:  0,
	},
}
