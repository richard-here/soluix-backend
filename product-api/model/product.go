package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"code"`
	Name        string         `json:"name" validate:"required"`
	Subcategory string         `json:"subcategory" validate:"required"`
	Brand       string         `json:"brand" validate:"required"`
	RetailPrice int            `json:"retail_price" validate:"required,min=0,max=50000000"`
	Status      bool           `json:"status" validate:"required"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
