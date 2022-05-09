package candidatevoterepo

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
func (r repositorySql) Create(entity models.CandidateVote) (*models.CandidateVote, error) {
	tx := r.db.Create(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &entity, nil
}

// Get all records
func (r repositorySql) Find() ([]models.CandidateVote, error) {
	entities := []models.CandidateVote{}

	tx := r.db.Find(&entities)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, tx.Error
	}

	return entities, nil
}

// Retrieving objects with primary key
func (r repositorySql) FindByCandidateId(candidateId int) (*models.CandidateVote, error) {
	var entity models.CandidateVote
	tx := r.db.Where("candidate_id = ?", candidateId).Find(&entity)
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
func (r repositorySql) IncreaseByCandidateId(candidateId int) (int, error) {
	tx := r.db.Model(&models.CandidateVote{}).Where("candidate_id = ?", candidateId).UpdateColumn("voted_count", gorm.Expr("voted_count + ?", 1))
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return int(tx.RowsAffected), tx.Error
	}

	return int(tx.RowsAffected), nil
}
