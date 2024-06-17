package position

import (
	"fmt"
	"time"

	"github.com/alsey89/people-matter/schema"
)

// ! Position ------------------------------------------------------------
func (d *Domain) GetPositions(companyID uint) ([]schema.Position, error) {
	db := d.params.Database.GetDB()

	var positions []schema.Position

	result := db.
		Where("company_id = ?", companyID).
		Find(&positions)
	if result.Error != nil {
		return nil, fmt.Errorf("[GetPositions] %w", result.Error)
	}

	return positions, nil
}

// ! UserPosition ------------------------------------------------------------
func (d *Domain) CreateUserPosition(companyID uint, userID uint, userPosition *schema.UserPosition) error {
	db := d.params.Database.GetDB()

	userPosition.CompanyID = companyID
	userPosition.UserID = userID
	userPosition.StartDate = time.Now()

	result := db.Create(userPosition)
	if result.Error != nil {
		return fmt.Errorf("[CreateUserPosition] %w", result.Error)
	}

	return nil
}

func (d *Domain) EndUserPosition(companyID uint, userID uint, userPositionID uint) error {
	db := d.params.Database.GetDB()

	dataToUpdate := map[string]interface{}{
		"end_date": time.Now(),
	}

	result := db.Model(&schema.UserPosition{}).
		Where("company_id = ? AND user_id = ? AND id = ? AND end_date IS NULL", companyID, userID, userPositionID).
		Updates(dataToUpdate)
	if result.Error != nil {
		return fmt.Errorf("[EndUserPosition] %w", result.Error)
	}

	return nil
}
