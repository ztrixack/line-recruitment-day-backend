package mocks

import (
	"election-service/internal/core/models"
	"time"

	"gorm.io/gorm"
)

var NewVoter = models.Voter{
	Model: models.Model{
		Id:        0,
		CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		DeletedAt: gorm.DeletedAt{
			Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Valid: false,
		},
	},
	CandidateId: 1,
	NationalId:  "1234567890123",
}

var Voter = models.Voter{
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
	NationalId:  "1234567890123",
}
var NilVoter = models.Voter{
	Model: models.Model{
		Id:        0,
		CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		DeletedAt: gorm.DeletedAt{
			Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Valid: false,
		},
	},
	CandidateId: 0,
	NationalId:  "",
}
