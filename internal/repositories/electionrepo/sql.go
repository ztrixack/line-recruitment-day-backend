package electionrepo

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

// Retrieving objects with primary key
func (r repositorySql) FindById(id int) (*models.Election, error) {
	var entity models.Election
	tx := r.db.Find(&entity, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, tx.Error
	}

	if tx.RowsAffected != 1 {
		return nil, nil
	}

	return &entity, nil
}

// Update with primary key
func (r repositorySql) UpdateById(id int, entity models.Json) (int, error) {
	tx := r.db.Model(&models.Election{}).Where("id = ?", id).Updates(&entity)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return int(tx.RowsAffected), tx.Error
	}

	return int(tx.RowsAffected), nil
}
