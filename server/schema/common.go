package schema

import (
	"time"
)

// ------------------------------------------------------------------------------------------------
type BaseModel struct {
	ID uint `json:"id" gorm:"primaryKey"`

	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt *time.Time `gorm:"index"`
}

// ------------------------------------------------------------------------------------------------
type ApprovalStatusEnum string

const (
	Pending  ApprovalStatusEnum = "pending"
	Approved ApprovalStatusEnum = "approved"
	Rejected ApprovalStatusEnum = "rejected"
)
