package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModels struct to generate model id
type BaseModels struct {
	Id       string `json:"id" gorm:"primaryKey;unique"`
	CreateBy string `gorm:"column:create_by"`
	UpdateBy string `gorm:"column:update_by"`
}

// BaseCUModels struct to generate CreatedAt, UpdatedAt
type BaseCUModels struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// BeforeCreate create uuid before model create
func (base *BaseModels) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	base.Id = uuid.String()
	return nil
}
