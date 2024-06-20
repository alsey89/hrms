package company

import (
	"errors"
	"fmt"
	"log"

	"github.com/alsey89/people-matter/schema"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//! Company ------------------------------------------------------------

// Creates new company, new root user, and create root role if it doesn't exist
// assigns root role and company to the new root user
func (d *Domain) CreateNewCompanyAndRootUser(incomingData *NewCompany) (*uint, *schema.User, error) {
	db := d.params.Database.GetDB()

	// Create company
	newCompany := schema.Company{
		Name:        incomingData.CompanyName,
		CompanySize: incomingData.CompanySize,
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(incomingData.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, fmt.Errorf("[CreateNewCompanyAndRootUser] Error hashing password: %w", err)
	}

	// Create root user
	newRootUser := schema.User{
		Email:    incomingData.RootUserEmail,
		Password: string(hashedPassword),
		IsActive: false,
	}

	// *----- Transaction Start -----
	tx := db.Begin()

	// Create company
	if err := tx.Create(&newCompany).Error; err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("[CreateNewCompanyAndRootUser] Error creating company: %w", err)
	}

	//* Create root role if it doesn't exist
	var rootRole schema.Role
	err = tx.
		Where("level = ? AND company_id = ?", "root", newCompany.ID).
		First(&rootRole).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		rootRole = schema.Role{
			CompanyID:   newCompany.ID,
			Level:       "root",
			Description: "Root User with full access to the company resources",
		}
		if err := tx.Create(&rootRole).Error; err != nil {
			tx.Rollback()
			return nil, nil, fmt.Errorf("[CreateNewCompanyAndRootUser] Error creating admin role: %w", err)
		}
	}

	// Create user
	err = tx.Create(&newRootUser).Error
	if err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("[CreateNewCompanyAndRootUser] Error creating user: %w", err)
	}

	// Assign root role to the root user
	err = tx.Model(&newRootUser).Association("Role").Append(&rootRole)
	if err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("[CreateNewCompanyAndRootUser] Error assigning root role to user: %w", err)
	}

	// Assign the new company to the new root user
	err = tx.Model(&newRootUser).Association("Companies").Append(&newCompany)
	if err != nil {
		tx.Rollback()
		return nil, nil, fmt.Errorf("[CreateNewCompanyAndRootUser] Error assigning company to user: %w", err)
	}

	tx.Commit()
	// *----- Transaction End -----

	return &newCompany.ID, &newRootUser, nil
}

// Get company data without preloading data
func (d *Domain) GetCompany(companyID *uint) (*schema.Company, error) {
	db := d.params.Database.GetDB()

	var existingCompany schema.Company

	result := db.First(&existingCompany, companyID)
	if result.Error != nil {
		return nil, fmt.Errorf("[GetCompany] %w", result.Error)
	}

	return &existingCompany, nil
}

// Gets company data, preloads departments, locations, and positions.
// Todo: benchmark and compare against joins
func (d *Domain) GetCompanyWithDetails(companyID *uint) (*schema.Company, error) {
	db := d.params.Database.GetDB()

	var existingCompany schema.Company

	result := db.
		Preload("Departments").
		Preload("Locations").
		Preload("Positions").
		First(&existingCompany, companyID)
	if result.Error != nil {
		return nil, fmt.Errorf("[GetCompanyAndExpand] %w", result.Error)
	}

	return &existingCompany, nil
}

// Updates company data, allows null values
func (d *Domain) UpdateCompany(companyID *uint, newData *schema.Company) error {
	db := d.params.Database.GetDB()

	//using a map instead of struct allows for null values
	dataToUpdate := map[string]interface{}{
		"Name":    newData.Name,
		"LogoURL": newData.LogoURL,
		"Website": newData.Website,
		"Email":   newData.Email,

		"Phone":       newData.Phone,
		"Address":     newData.Address,
		"City":        newData.City,
		"State":       newData.State,
		"Country":     newData.Country,
		"PostalCode":  newData.PostalCode,
		"CompanySize": newData.CompanySize,
	}

	result := db.
		Model(&schema.Company{}).
		Where("id = ?", companyID).
		Updates(dataToUpdate)
	if result.Error != nil {
		return fmt.Errorf("[UpdateCompany] %w", result.Error)
	}

	return nil
}

// Delete company data
func (d *Domain) DeleteCompany(companyID *uint) error {
	db := d.params.Database.GetDB()

	result := db.Delete(&schema.Company{}, companyID)
	if result.Error != nil {
		return fmt.Errorf("[DeleteCompany] %w", result.Error)
	}

	return nil
}

//! Department ------------------------------------------------------------

// Create new department
func (d *Domain) CreateDepartment(newDepartment *schema.Department) error {
	db := d.params.Database.GetDB()

	result := db.Create(newDepartment)
	if result.Error != nil {
		return fmt.Errorf("[CreateDepartment] %w", result.Error)
	}

	return nil
}

// Update department data, allows null values
func (d *Domain) UpdateDepartment(companyID *uint, departmentID *uint, newData *schema.Department) error {
	db := d.params.Database.GetDB()

	dataToUpdate := map[string]interface{}{
		"Name":        newData.Name,
		"Description": newData.Description,
	}

	result := db.
		Model(&schema.Department{}).
		Where("id = ? AND company_id = ?", departmentID, companyID).
		Updates(dataToUpdate)
	if result.Error != nil {
		return fmt.Errorf("[UpdateDepartment] %w", result.Error)
	}

	return nil
}

func (d *Domain) DeleteDepartment(companyID *uint, departmentID *uint) error {
	db := d.params.Database.GetDB()

	result := db.
		Model(&schema.Department{}).
		Where("company_id = ? AND id = ?", companyID, departmentID).
		Delete(&schema.Department{})
	if result.Error != nil {
		return fmt.Errorf("[DeleteDepartment] %w", result.Error)
	}

	return nil
}

//! Location ------------------------------------------------------------

// Create new location
func (d *Domain) CreateLocation(newLocation *schema.Location) error {
	db := d.params.Database.GetDB()

	result := db.Create(newLocation)
	if result.Error != nil {
		return fmt.Errorf("[CreateLocation] %w", result.Error)
	}

	return nil
}

// Update all location data, allows null values
func (d *Domain) UpdateLocation(companyID *uint, locationID *uint, newData *schema.Location) error {
	log.Printf("UpdateLocation: %v", newData)
	db := d.params.Database.GetDB()

	dataToUpdate := map[string]interface{}{
		"Name":         newData.Name,
		"IsHeadOffice": newData.IsHeadOffice,
		"Phone":        newData.Phone,
		"Address":      newData.Address,
		"City":         newData.City,
		"State":        newData.State,
		"Country":      newData.Country,
		"PostalCode":   newData.PostalCode,
	}

	result := db.
		Model(&schema.Location{}).
		Where("company_id = ? AND id = ?", companyID, locationID).
		Updates(dataToUpdate)
	if result.Error != nil {
		return fmt.Errorf("[UpdateLocation] %w", result.Error)
	}

	return nil
}

// Update location data, no HeadOffice value
func (d *Domain) UpdateLocationNoHeadOffice(companyID *uint, locationID *uint, newData *schema.Location) error {

	db := d.params.Database.GetDB()

	dataToUpdate := map[string]interface{}{
		"Name":       newData.Name,
		"Address":    newData.Address,
		"City":       newData.City,
		"State":      newData.State,
		"Country":    newData.Country,
		"PostalCode": newData.PostalCode,
	}

	result := db.
		Model(&schema.Location{}).
		Where("company_id = ? AND id = ?", companyID, locationID).
		Updates(dataToUpdate)
	if result.Error != nil {
		return fmt.Errorf("[UpdateLocation] %w", result.Error)
	}

	return nil
}

func (d *Domain) DeleteLocation(companyID *uint, locationID *uint) error {
	db := d.params.Database.GetDB()

	result := db.
		Model(&schema.Location{}).
		Where("company_id = ? AND id = ?", companyID, locationID).
		Delete(&schema.Location{})
	if result.Error != nil {
		return fmt.Errorf("[DeleteLocation] %w", result.Error)
	}

	return nil
}

//! Position ------------------------------------------------------------

func (d *Domain) CreatePosition(newPosition *schema.Position) error {
	db := d.params.Database.GetDB()

	result := db.Create(newPosition)
	if result.Error != nil {
		return fmt.Errorf("[CreatePosition] %w", result.Error)
	}

	return nil
}

func (d *Domain) UpdatePosition(companyID *uint, positionID *uint, newData *schema.Position) error {
	db := d.params.Database.GetDB()

	dataToUpdate := map[string]interface{}{
		"Title":          newData.Name,
		"Description":    newData.Description,
		"Duties":         newData.Duties,
		"Qualifications": newData.Qualifications,
		"Experience":     newData.Experience,
		"MinSalary":      newData.MinSalary,
		"MaxSalary":      newData.MaxSalary,
		"DepartmentID":   newData.DepartmentID,
		"ManagerID":      newData.ManagerID,
	}

	result := db.
		Model(&schema.Position{}).
		Where("company_id = ? AND id = ?", companyID, positionID).
		Updates(dataToUpdate)
	if result.Error != nil {
		return fmt.Errorf("[UpdatePosition] %w", result.Error)
	}

	return nil
}

func (d *Domain) DeletePosition(companyID *uint, positionID *uint) error {
	db := d.params.Database.GetDB()

	result := db.
		Model(&schema.Position{}).
		Where("company_id = ? AND id = ?", companyID, positionID).
		Delete(&schema.Position{})
	if result.Error != nil {
		return fmt.Errorf("[DeletePosition] %w", result.Error)
	}

	return nil
}
