package schema

import (
	"time"
)

// User Schema ---------------------------------------------------------------
type User struct {
	BaseModel
	UserRoles []UserRole `json:"userRoles" gorm:"foreignKey:UserID"`
	IsActive  bool       `json:"isActive" gorm:"default:false"`
	// ------------------------------------------------------------------------------------------------
	Email      string      `json:"email" gorm:"uniqueIndex;type:varchar(100);not null"`
	Password   string      `json:"-" gorm:"type:varchar(100)"` //* Password is not returned in JSON
	AvatarURL  string      `json:"avatarUrl" gorm:"type:text"`
	LastLogin  *time.Time  `json:"lastLogin" gorm:"default:null"`
	IsArchived bool        `json:"isArchived" gorm:"default:false"`
	Documents  []*Document `json:"documents" gorm:"foreignKey:UserID"`
	// ------------------------------------------------------------------------------------------------
	FirstName        string            `json:"firstName"`
	MiddleName       *string           `json:"middleName"`
	LastName         string            `json:"lastName"`
	Nickname         string            `json:"nickname"`
	ContactInfo      *ContactInfo      `json:"contactInfo" gorm:"foreignKey:UserID"`
	EmergencyContact *EmergencyContact `json:"emergencyContact" gorm:"foreignKey:UserID"`
	// ------------------------------------------------------------------------------------------------
	UserPositions []UserPosition `json:"userPositions" gorm:"foreignKey:UserID"`
	// ------------------------------------------------------------------------------------------------
	SalaryID *uint     `json:"salaryId" gorm:"foreignKey:UserID"`
	Payments []Payment `json:"payments" gorm:"foreignKey:UserID"`
	// ------------------------------------------------------------------------------------------------
	Leave      []Leave      `json:"leave" gorm:"foreignKey:UserID"`
	Attendance []Attendance `json:"attendance" gorm:"foreignKey:UserID"`
}

type UserRole struct {
	BaseModel
	User      *User    `json:"user" gorm:"constraint:OnDelete:CASCADE;foreignKey:UserID"`
	UserID    uint     `json:"userId" gorm:"constraint:OnDelete:CASCADE;not null"`
	Role      *Role    `json:"role" gorm:"constraint:OnDelete:CASCADE;foreignKey:RoleID"`
	RoleID    uint     `json:"roleId" gorm:"constraint:OnDelete:CASCADE;not null"`
	Company   *Company `json:"company" gorm:"constraint:OnDelete:CASCADE;foreignKey:CompanyID"`
	CompanyID uint     `json:"companyId" gorm:"constraint:OnDelete:CASCADE;not null"`
	// optionally scope to location
	Location   *Location `json:"location" gorm:"constraint:OnDelete:SET NULL;foreignKey:LocationID"`
	LocationID *uint     `json:"locationId" gorm:"constraint:OnDelete:CASCADE"`
}

// Role Schema ---------------------------------------------------------------

type Role struct {
	BaseModel
	CompanyID   uint   `json:"companyId" gorm:"constraint:OnDelete:CASCADE;not null"`
	Name        string `json:"name" gorm:"type:varchar(100);not null"`
	Description string `json:"description"`
}

// ContactInfo Schema --------------------------------------------------------

type ContactInfo struct {
	BaseModel
	CompanyID uint `json:"companyId" gorm:"constraint:OnDelete:CASCADE;not null"`
	// ------------------------------------------------------------------------------------------------
	UserID     uint   `json:"userId" gorm:"constraint:OnDelete:CASCADE;not null"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	// ------------------------------------------------------------------------------------------------
}

// EmergencyContact Schema ---------------------------------------------------

type EmergencyContact struct {
	BaseModel
	CompanyID uint `json:"companyId" gorm:"constraint:OnDelete:CASCADE;not null"`
	// ------------------------------------------------------------------------------------------------
	UserID     uint    `json:"userId" gorm:"constraint:OnDelete:CASCADE;not null"`
	FirstName  string  `json:"firstName"`
	MiddleName *string `json:"middleName"`
	LastName   string  `json:"lastName"`
	Relation   string  `json:"relation"`
	Mobile     string  `json:"mobile"`
	Email      string  `json:"email"`
	// ------------------------------------------------------------------------------------------------
}
