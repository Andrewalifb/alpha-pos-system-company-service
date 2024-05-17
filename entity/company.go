package entity

import (
	"time"

	"github.com/google/uuid"
)

type PosCompany struct {
	CompanyID   uuid.UUID `gorm:"type:uuid;primary_key" json:"company_id"`
	CompanyName string    `gorm:"type:varchar(255);not null" json:"company_name"`
	CreatedAt   time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy   uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt   time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy   uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
