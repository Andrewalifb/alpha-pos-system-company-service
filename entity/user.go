package entity

import (
	"time"

	"github.com/google/uuid"
)

type PosUser struct {
	UserID       uuid.UUID  `gorm:"type:uuid;primary_key" json:"user_id"`
	Username     string     `gorm:"type:varchar(255);not null" json:"username"`
	PasswordHash string     `gorm:"type:varchar(255);not null" json:"password_hash"`
	RoleID       uuid.UUID  `gorm:"type:uuid;not null" json:"role_id"`
	CompanyID    *uuid.UUID `gorm:"type:uuid" json:"company_id"`
	BranchID     *uuid.UUID `gorm:"type:uuid" json:"branch_id"`
	StoreID      *uuid.UUID `gorm:"type:uuid" json:"store_id"`
	FirstName    string     `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName     string     `gorm:"type:varchar(255);not null" json:"last_name"`
	Email        string     `gorm:"type:varchar(255)" json:"email"`
	PhoneNumber  string     `gorm:"type:varchar(20)" json:"phone_number"`
	CreatedAt    time.Time  `gorm:"type:timestamp" json:"created_at"`
	CreatedBy    uuid.UUID  `gorm:"type:uuid" json:"created_by"`
	UpdatedAt    time.Time  `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy    uuid.UUID  `gorm:"type:uuid" json:"updated_by"`
}
