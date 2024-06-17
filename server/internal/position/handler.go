package position

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/alsey89/people-matter/internal/common"
	"github.com/alsey89/people-matter/schema"
)

// ! Admin -----------------------------------------------------------
func (d *Domain) GetAllPositionsHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		return d.handleError(c, "error getting company id from token", err)
	}
	if companyID == nil {
		return d.handleError(c, "company ID is nil", nil)
	}

	positions, err := d.GetPositions(*companyID)
	if err != nil {
		return d.handleError(c, "error getting positions", err)
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "positions retrieved",
		Data:    positions,
	})
}

// Creates a new user position, gets companyID from token, userID from param, and userPosition from body
func (d *Domain) AssignPositionHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		return d.handleError(c, "error getting company id from token", err)
	}
	if companyID == nil {
		return d.handleError(c, "company ID is nil", nil)
	}

	userID, err := common.GetIDFromParam("userId", c)
	if err != nil {
		return d.handleError(c, "error getting user id from param", err)
	}
	if userID == nil {
		return d.handleError(c, "user ID is nil", nil)
	}

	userPosition := new(schema.UserPosition)
	if err := c.Bind(userPosition); err != nil {
		return d.handleError(c, "error binding user position", err)
	}

	// Create user position
	if err := d.CreateUserPosition(*companyID, *userID, userPosition); err != nil {
		return d.handleError(c, "error creating user position", err)
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "user position created",
		Data:    nil,
	})
}

// Updates a user position, gets companyID from token, userID from param, and userPosition from body
func (d *Domain) UnassignPositionHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		return d.handleError(c, "error getting company id from token", err)
	}
	if companyID == nil {
		return d.handleError(c, "company ID is nil", nil)
	}

	userID, err := common.GetIDFromParam("userId", c)
	if err != nil {
		return d.handleError(c, "error getting user id from param", err)
	}
	if userID == nil {
		return d.handleError(c, "user ID is nil", nil)
	}

	userPositionID, err := common.GetIDFromParam("userPositionId", c)
	if err != nil {
		return d.handleError(c, "error getting user position id from param", err)
	}
	if userPositionID == nil {
		return d.handleError(c, "user position ID is nil", nil)
	}

	// Update user position
	if err := d.EndUserPosition(*companyID, *userID, *userPositionID); err != nil {
		return d.handleError(c, "error updating user position", err)
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "user position updated",
		Data:    nil,
	})
}

// Helper method for error handling
func (d *Domain) handleError(c echo.Context, message string, err error) error {
	if err != nil {
		d.logger.Error(message, zap.Error(err))
	} else {
		d.logger.Error(message)
	}
	return c.JSON(http.StatusInternalServerError, common.APIResponse{
		Message: message,
		Data:    nil,
	})
}
