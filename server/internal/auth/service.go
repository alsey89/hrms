package auth

import (
	"fmt"

	"github.com/alsey89/people-matter/schema"
	"golang.org/x/crypto/bcrypt"
)

func (d *Domain) AuthenticateUserService(email, password string) (*schema.User, []schema.Company, error) {
	db := d.params.Database.GetDB()
	var user schema.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, nil, err
	}

	// Check if user is active
	if !user.IsActive {
		return nil, nil, fmt.Errorf("[SignIn]: %w", ErrUserNotConfirmed)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, nil, ErrInvalidCredentials
	}

	var companies []schema.Company
	err = db.Model(&user).Association("Companies").Find(&companies)
	if err != nil {
		return nil, nil, err
	}

	return &user, companies, nil
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

func (d *Domain) ConfirmEmailService(userID uint, companyID uint) error {
	db := d.params.Database.GetDB()

	//query user and mark as confirmed
	result := db.Model(&schema.User{}).
		Where("id = ?", userID).
		Where("id IN (SELECT user_id FROM user_companies WHERE company_id = ?)", companyID).
		Update("is_active", true)
	if result.Error != nil {
		return fmt.Errorf("[Confirmation]: %w", result.Error)
	}

	return nil
}
