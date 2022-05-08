package models

import (
	"time"

	"gorm.io/gorm"
)

type Json = map[string]interface{}

type Model struct {
	Id        uint           `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
