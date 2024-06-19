package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/alsey89/people-matter/internal/common"
	"github.com/alsey89/people-matter/schema"
)

// ! Manager ---------------------------------------------------------
func (d *Domain) GetAllLocationUsersHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	locationID, err := common.GetLocationIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting location id from token",
			Data:    nil,
		})
	}

	users, err := d.GetUsersByLocation(companyID, locationID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting location users",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "success",
		Data:    users,
	})
}
func (d *Domain) CreateLocationUserHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	var user schema.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error binding user",
			Data:    nil,
		})
	}

	err = d.CreateUser(companyID, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error creating user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "success",
		Data:    nil,
	})
}
func (d *Domain) UpdateLocationUserHandler(c echo.Context) error {

	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	locationID, err := common.GetIDFromParam("location_id", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting location id from param",
			Data:    nil,
		})
	}
	if locationID == nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting location id from param",
			Data:    nil,
		})
	}

	userID, err := common.GetIDFromParam("userId", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting user id from param",
			Data:    nil,
		})
	}

	var user schema.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error binding user",
			Data:    nil,
		})
	}

	if err := d.UpdateUserBasicInformation(companyID, userID, &user); err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error updating user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "success",
		Data:    nil,
	})
}
func (d *Domain) DeleteLocationUserHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	locationID, err := common.GetLocationIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting location id from param",
			Data:    nil,
		})
	}
	if locationID == nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting location id from param",
			Data:    nil,
		})
	}

	userID, err := common.GetUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting user id from param",
			Data:    nil,
		})
	}

	if err := d.DeleteUser(companyID, userID); err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error deleting user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "success",
		Data:    nil,
	})
}

// ! Admin -----------------------------------------------------------
func (d *Domain) GetAllUsersHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[GetAllUsersHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	users, err := d.GetAllUsers(companyID)
	if err != nil {
		d.logger.Error("[GetAllUsersHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting users",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "success",
		Data:    users,
	})
}
func (d *Domain) CreateUserHandler(c echo.Context) error {

	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	user := new(schema.User)

	err = c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error binding user",
			Data:    nil,
		})
	}

	if err := d.CreateUser(companyID, user); err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error creating user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "user created",
		Data:    nil,
	})
}
func (d *Domain) UpdateUserHandler(c echo.Context) error {

	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[UpdateUserHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}
	if companyID == nil {
		d.logger.Error("[UpdateUserHandler] companyID is nil")
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	userID, err := common.GetIDFromParam("userId", c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting user id from param",
			Data:    nil,
		})
	}

	user := new(schema.User)

	err = c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error binding user",
			Data:    nil,
		})
	}

	err = d.UpdateUserBasicInformation(companyID, userID, user)
	if err != nil {
		d.logger.Error("[UpdateUserHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error updating user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "user updated",
		Data:    nil,
	})
}
func (d *Domain) DeleteUserHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[DeleteUserHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	userID, err := common.GetIDFromParam("userId", c)
	if err != nil {
		d.logger.Error("[DeleteUserHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting user id from param",
			Data:    nil,
		})
	}

	err = d.DeleteUser(companyID, userID)
	if err != nil {
		d.logger.Error("[DeleteUserHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error deleting user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "user deleted",
		Data:    nil,
	})
}
