package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomPublicModel struct {
	ID        uuid.UUID       `gorm:"type:uuid;column:id;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"-"`
}
