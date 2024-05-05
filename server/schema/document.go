package schema

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	CompanyID uint `json:"company_id" gorm:"onUpdate:CASCADE onDelete:CASCADE;not null"`
	// ------------------------------------------------------------------------------------------------
	UserID uint   `json:"userId"      gorm:"onUpdate:CASCADE;onDelete:CASCADE"`
	URL    string `json:"url"         gorm:"type:text;not null"`
	Notes  string `json:"description" gorm:"type:text"`
	// ------------------------------------------------------------------------------------------------
}
