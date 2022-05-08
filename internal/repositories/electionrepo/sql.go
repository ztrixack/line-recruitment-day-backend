package electionrepo

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

// Retrieving objects with primary key
func (r repositorySQL) FindByID(id int) (*models.Election, error) {
	var entity models.Election
	tx := r.db.Find(&entity, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, nil
	}

	return &entity, nil
}
