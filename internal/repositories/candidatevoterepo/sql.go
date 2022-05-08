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

// Get all records
func (r repositorySQL) Find() ([]models.CandidateVote, error) {
	entities := []models.CandidateVote{}

	tx := r.db.Find(&entities)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return entities, nil
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

// Update with primary key
func (r repositorySQL) UpdateByID(id int, entity models.Json) (int, error) {
	tx := r.db.Model(&models.CandidateVote{}).Where("id = ?", id).Updates(&entity)
	if tx.Error != nil {
		return int(tx.RowsAffected), tx.Error
	}

	return int(tx.RowsAffected), nil
}
