package entity

import (
	"time"

	"github.com/google/uuid"
)

type PosRole struct {
	RoleID    uuid.UUID `gorm:"type:uuid;primary_key" json:"role_id"`
	RoleName  string    `gorm:"type:varchar(50);not null" json:"role_name"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
