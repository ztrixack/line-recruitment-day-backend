package voterrepo

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

// Create Record
func (r repositorySql) Create(entity models.Voter) (*models.Voter, error) {
	tx := r.db.Create(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &entity, nil
}

// Retrieving object
func (r repositorySql) FindByNationId(nationId string) (*models.Voter, error) {
	var entity models.Voter
	tx := r.db.Model(&models.Voter{}).Where("national_id = ?", nationId).Find(&entity)
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
	tx := r.db.Model(&models.Voter{}).Where("id = ?", id).Updates(&entity)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return int(tx.RowsAffected), tx.Error
	}

	return int(tx.RowsAffected), nil
}
