package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"code"`
	Name        string         `json:"name"`
	Subcategory string         `json:"subcategory"`
	Brand       string         `json:"brand"`
	RetailPrice int            `json:"retail_price"`
	Status      bool           `json:"status"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
