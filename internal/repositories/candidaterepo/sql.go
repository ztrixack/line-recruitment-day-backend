package candidaterepo

import (
	"election-service/internal/core/models"

	"gorm.io/gorm"
)

type repositorySQL struct {
	db *gorm.DB
}

func NewSQL(db *gorm.DB) repositorySQL {
	return repositorySQL{db: db}
}

// Get all records
func (r repositorySQL) Find() ([]models.Candidate, error) {
	entities := []models.Candidate{}

	tx := r.db.Find(&entities)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return entities, nil
}
