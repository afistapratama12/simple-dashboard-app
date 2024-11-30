package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID string `gorm:"column:id;primarykey;type:char(36)" json:"id"`
	BaseModelTime
}

type BaseModelTime struct {
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP;" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}

func (b *BaseModel) SetBaseData() (isEmptyID bool) {
	if b.ID == "" {
		isEmptyID = true
		b.ID = uuid.NewString()
	}
	b.UpdatedAt = time.Now()
	if b.CreatedAt.IsZero() {
		b.CreatedAt = time.Now()
	}
	return
}
