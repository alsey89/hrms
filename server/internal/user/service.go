package user

import (
	"fmt"
	"time"

	"github.com/alsey89/people-matter/schema"
)

//! User ------------------------------------------------------------

// Get all users in company
func (d *Domain) GetAllUsers(companyID *uint) ([]schema.User, error) {
	db := d.params.Database.GetDB()

	var users []schema.User

	result := db.
		Preload("UserPositions", "end_date IS NULL OR end_date > ?", time.Now()). // Only get active user positions
		Where("company_id = ?", companyID).
		Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("[GetAllUsers] %w", result.Error)
	}

	return users, nil
}

// Get all users in company and location
func (d *Domain) GetUsersByLocation(companyID, locationID *uint) ([]schema.User, error) {
	db := d.params.Database.GetDB()

	var users []schema.User

	result := db.
		Joins("JOIN user_positions ON user_positions.user_id = users.id").
		Where("users.company_id = ? AND user_positions.location_id = ?", companyID, locationID).
		Preload("UserPositions", "location_id = ? AND (end_date IS NULL OR end_date > ?)", locationID, time.Now()).
		Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("[GetAllUsers] %w", result.Error)
	}

	return users, nil
}

// Get a single user by ID
func (d *Domain) GetUserByID(companyID *uint, userID *uint) (*schema.User, error) {
	db := d.params.Database.GetDB()

	var existingUser schema.User

	result := db.
		Preload("ContactInfo").
		Preload("EmergencyContact").
		Where("company_id = ? AND id = ?", companyID, userID).
		First(&existingUser)
	if result.Error != nil {
		return nil, fmt.Errorf("[GetUser] %w", result.Error)
	}

	return &existingUser, nil
}

// Create new user
func (d *Domain) CreateUser(newUser *schema.User) error {
	db := d.params.Database.GetDB()

	result := db.Create(newUser)
	if result.Error != nil {
		return fmt.Errorf("[CreateUser] %w", result.Error)
	}

	return nil
}

// Update user data, allows null values
func (d *Domain) UpdateUserBasicInformation(companyID *uint, userID *uint, newData *schema.User) error {
	db := d.params.Database.GetDB()

	dataToUpdate := map[string]interface{}{
		"AvatarURL":  newData.AvatarURL,
		"FirstName":  newData.FirstName,
		"MiddleName": newData.MiddleName,
		"LastName":   newData.LastName,
		"Nickname":   newData.Nickname,
	}

	result := db.
		Model(&schema.User{}).
		Where("company_id = ? AND id = ?", companyID, userID).
		Updates(dataToUpdate)
	if result.Error != nil {
		return fmt.Errorf("[UpdateUser] %w", result.Error)
	}

	return nil
}

// Update user contact information
func (d *Domain) UpdateUserContactInformation(companyID *uint, userID *uint, newData *schema.ContactInfo) error {
	db := d.params.Database.GetDB()

	dataToUpdate := map[string]interface{}{
		"Address":    newData.Address,
		"City":       newData.City,
		"State":      newData.State,
		"PostalCode": newData.PostalCode,
		"Country":    newData.Country,
		"Mobile":     newData.Mobile,
		"Email":      newData.Email,
	}

	result := db.
		Model(&schema.ContactInfo{}).
		Where("company_id = ? AND user_id = ?", companyID, userID).
		Updates(dataToUpdate)
	if result.Error != nil {
		return fmt.Errorf("[UpdateUserContactInformation] %w", result.Error)
	}

	return nil
}

// Update user emergency contact
func (d *Domain) UpdateUserEmergencyContact(companyID *uint, userID *uint, newData *schema.EmergencyContact) error {

	db := d.params.Database.GetDB()

	dataToUpdate := map[string]interface{}{
		"FirstName": newData.FirstName,
		"LastName":  newData.LastName,
		"Mobile":    newData.Mobile,
		"Email":     newData.Email,
	}

	result := db.
		Model(&schema.EmergencyContact{}).
		Where("company_id = ? AND user_id = ?", companyID, userID).
		Updates(dataToUpdate)
	if result.Error != nil {
		return fmt.Errorf("[UpdateUserEmergencyContact] %w", result.Error)
	}

	return nil
}

// Delete user data
func (d *Domain) DeleteUser(companyID *uint, userID *uint) error {
	db := d.params.Database.GetDB()

	result := db.
		Model(&schema.User{}).
		Where("company_id = ? AND id = ?", companyID, userID).
		Delete(&schema.User{})
	if result.Error != nil {
		return fmt.Errorf("[DeleteUser] %w", result.Error)
	}

	return nil
}
