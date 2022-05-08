package mocks

import (
	"election-service/internal/core/models"
	"time"

	"gorm.io/gorm"
)

var Candidates = []models.Candidate{
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
		Name:      "Elon Mask",
		Dob:       "June 28, 1971",
		BioLink:   "<https://en.wikipedia.org/wiki/Elon_Musk>",
		ImageLink: "<https://upload.wikipedia.org/wikipedia/commons/e/ed/Elon_Musk_Royal_Society.jpg>",
		Policy:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
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
		Name:      "Brown",
		Dob:       "August 8, 2011",
		BioLink:   "<https://line.fandom.com/wiki/Brown>",
		ImageLink: "<https://static.wikia.nocookie.net/line/images/b/bb/2015-brown.png/revision/latest/scale-to-width-down/700?cb=20150808131630>",
		Policy:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
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
		Name:      "LINE Brown",
		Dob:       "August 8, 2011",
		BioLink:   "<https://line.fandom.com/wiki/Brown>",
		ImageLink: "<https://static.wikia.nocookie.net/line/images/b/bb/2015-brown.png/revision/latest/scale-to-width-down/700?cb=20150808131630>",
		Policy:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
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
		Name:      "Jeff Bezos",
		Dob:       "January 12, 1964",
		BioLink:   "<https://en.wikipedia.org/wiki/Jeff_Bezos>",
		ImageLink: "<https://pbs.twimg.com/profile_images/669103856106668033/UF3cgUk4_400x400.jpg>",
		Policy:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
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
		Name:      "Elon Mask",
		Dob:       "June 28, 1971",
		BioLink:   "<https://en.wikipedia.org/wiki/Elon_Musk>",
		ImageLink: "<https://upload.wikipedia.org/wikipedia/commons/e/ed/Elon_Musk_Royal_Society.jpg>",
		Policy:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
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
		Name:      "Brown",
		Dob:       "August 8, 2011",
		BioLink:   "<https://line.fandom.com/wiki/Brown>",
		ImageLink: "<https://static.wikia.nocookie.net/line/images/b/bb/2015-brown.png/revision/latest/scale-to-width-down/700?cb=20150808131630>",
		Policy:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
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
		Name:      "LINE Brown",
		Dob:       "August 8, 2011",
		BioLink:   "<https://line.fandom.com/wiki/Brown>",
		ImageLink: "<https://static.wikia.nocookie.net/line/images/b/bb/2015-brown.png/revision/latest/scale-to-width-down/700?cb=20150808131630>",
		Policy:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
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
		Name:      "Jeff Bezos",
		Dob:       "January 12, 1964",
		BioLink:   "<https://en.wikipedia.org/wiki/Jeff_Bezos>",
		ImageLink: "<https://pbs.twimg.com/profile_images/669103856106668033/UF3cgUk4_400x400.jpg>",
		Policy:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown",
	},
}
