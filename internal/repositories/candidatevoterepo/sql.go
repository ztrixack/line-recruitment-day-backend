package candidatevoterepo

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

// Create Record
func (r repositorySQL) Create(entity models.CandidateVote) (*models.CandidateVote, error) {
	tx := r.db.Create(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &entity, nil
}

// Retrieving objects with primary key
func (r repositorySQL) FindByID(id int) (*models.CandidateVote, error) {
	var entity models.CandidateVote
	tx := r.db.Find(&entity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, nil
	}

	return &entity, nil
}
