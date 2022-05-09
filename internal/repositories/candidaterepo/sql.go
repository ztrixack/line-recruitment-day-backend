package candidaterepo

import (
	"election-service/internal/core/models"

	"gorm.io/gorm"
)

type repositorySql struct {
	db *gorm.DB
}

func NewSql(db *gorm.DB) repositorySql {
	return repositorySql{db: db}
}

// Get all records
func (r repositorySql) Find() ([]models.Candidate, error) {
	entities := []models.Candidate{}

	tx := r.db.Find(&entities)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, tx.Error
	}

	return entities, nil
}
