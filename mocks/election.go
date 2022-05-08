package mocks

import (
	"election-service/internal/core/models"
	"time"

	"gorm.io/gorm"
)

var SolicitElection = models.Election{
	Model: models.Model{
		Id:        1,
		CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		DeletedAt: gorm.DeletedAt{
			Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Valid: false,
		},
	},
	Enable: false,
	State:  "solicit",
}

var VotingElection = models.Election{
	Model: models.Model{
		Id:        1,
		CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		DeletedAt: gorm.DeletedAt{
			Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Valid: false,
		},
	},
	Enable: true,
	State:  "voting",
}

var ClosedElection = models.Election{
	Model: models.Model{
		Id:        1,
		CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		DeletedAt: gorm.DeletedAt{
			Time:  time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
			Valid: false,
		},
	},
	Enable: false,
	State:  "closed",
}

var UpdateToVoting = models.UpdateElectionData{
	Enable: true,
	State:  "voting",
}

var UpdateToClosed = models.UpdateElectionData{
	Enable: false,
	State:  "closed",
}
