package auth

import (
	"fmt"

	"github.com/alsey89/people-matter/schema"
	"golang.org/x/crypto/bcrypt"
)

func (d *Domain) AuthenticateUserService(email, password string) ([]schema.UserRole, error) {
	db := d.params.Database.GetDB()
	var user schema.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	// Check if user is active
	if !user.IsActive {
		return nil, fmt.Errorf("[SignIn]: %w", ErrUserNotConfirmed)
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Retrieve UserRoles with Company and Location information
	var userRoles []schema.UserRole
	err = db.
		Preload("Role").
		Preload("Company").
		Preload("Location").
		Where("user_id = ?", user.ID).Find(&userRoles).Error
	if err != nil {
		return nil, fmt.Errorf("[AuthenticateUserService]: %w", err)
	}

	return userRoles, nil
}

func (d *Domain) GetUserBySelectedUserRole(userRoleID *uint) (*schema.UserRole, error) {
	var preloadedUserRole schema.UserRole
	db := d.params.Database.GetDB()

	err := db.
		Preload("User").
		Preload("Role").
		Preload("Company").
		Preload("Location").
		Where("id = ?", userRoleID).
		First(&preloadedUserRole).Error
	if err != nil {
		return nil, fmt.Errorf("[GetUserBySelectedUserRole]: %w", err)
	}

	return &preloadedUserRole, nil
}

func (d *Domain) GetUserByEmailAndCompany(email string, companyId uint) (*schema.User, error) {
	var user schema.User
	db := d.params.Database.GetDB()
	err := db.
		Where("email = ?", email).
		Where("id IN (SELECT user_id FROM user_companies WHERE company_id = ?)", companyId).
		Preload("Role").
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *Domain) ConfirmEmailService(userID uint) error {
	db := d.params.Database.GetDB()

	//query user and mark as confirmed
	result := db.Model(&schema.User{}).
		Where("id = ?", userID).
		Update("is_active", true)
	if result.Error != nil {
		return fmt.Errorf("[Confirmation]: %w", result.Error)
	}

	return nil
}
