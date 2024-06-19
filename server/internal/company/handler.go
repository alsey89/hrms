package company

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/alsey89/people-matter/internal/common"
	"github.com/alsey89/people-matter/schema"
)

// ! Company ------------------------------------------------------------

// fetches company data *with* preloaded department, location, position data
func (d *Domain) GetCompanyHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[GetCompanyHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	companyData, err := d.GetCompanyWithDetails(companyID)
	if err != nil {
		d.logger.Error("[GetCompanyHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company data",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "detailed company data retrieved",
		Data:    companyData,
	})
}

// create company and admin user
func (d *Domain) CreateCompanyHandler(c echo.Context) error {

	form := new(NewCompany)
	err := c.Bind(form)
	if err != nil {
		d.logger.Error("[signupHandler] error binding credentials", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "form error",
			Data:    nil,
		})
	}

	// validate email
	email := form.RootUserEmail
	if !common.EmailValidator(email) {
		d.logger.Error("[signupHandler] email validation failed")
		return c.JSON(http.StatusBadRequest, common.APIResponse{
			Message: "invalid email",
			Data:    nil,
		})
	}

	// validate password
	password := form.Password
	confirmPassword := form.ConfirmPassword
	if password != confirmPassword {
		d.logger.Error("[signupHandler] password confirmation check failed")
		return c.JSON(http.StatusBadRequest, common.APIResponse{
			Message: "passwords do not match",
			Data:    nil,
		})
	}

	// validate company name
	companyName := form.CompanyName
	if companyName == "" {
		d.logger.Error("[signupHandler] company name is empty")
		return c.JSON(http.StatusBadRequest, common.APIResponse{
			Message: "company name required",
			Data:    nil,
		})
	}

	createdCompanyID, createdAdminUser, err := d.CreateNewCompanyAndRootUser(form)
	if err != nil {
		d.logger.Error("[createCompanyHandler]", zap.Error(err))
		if errors.Is(err, ErrUserExists) {
			return c.JSON(http.StatusConflict, common.APIResponse{
				Message: "user already exists",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error creating company and admin user",
			Data:    nil,
		})
	}
	if createdCompanyID == nil {
		d.logger.Error("[createCompanyHandler] created company id is nil")
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error creating company and admin user",
			Data:    nil,
		})
	}

	err = d.params.Auth.SendConfirmationEmail(createdAdminUser.Email, createdAdminUser.ID, *createdCompanyID)
	if err != nil {
		d.logger.Error("[createCompanyHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error sending confirmation email",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "company and admin user created",
		Data:    nil,
	})
}

// updates company data, gets companyID from token
func (d *Domain) UpdateCompanyHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[UpdateCompanyHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	dataToUpdate := new(schema.Company)

	err = c.Bind(dataToUpdate)
	if err != nil {
		d.logger.Error("[UpdateCompanyHandler] error binding dataToUpdate", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	err = d.UpdateCompany(companyID, dataToUpdate)
	if err != nil {
		d.logger.Error("[UpdateCompanyHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error updating company",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "company data updated",
		Data:    nil,
	})
}

// deletes company data, gets companyID from token
func (d *Domain) DeleteCompanyHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[DeleteCompanyHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	err = d.DeleteCompany(companyID)
	if err != nil {
		d.logger.Error("[DeleteCompanyHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error deleting company",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "company deleted",
		Data:    nil,
	})
}

//! Department ------------------------------------------------------------

// creates department, gets companyID from token
func (d *Domain) CreateDepartmentHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[CreateDepartmentHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}
	if companyID == nil {
		d.logger.Error("[CreateDepartmentHandler] companyID is nil")
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "company id is nil",
			Data:    nil,
		})
	}

	newDepartment := new(schema.Department)

	err = c.Bind(newDepartment)
	if err != nil {
		d.logger.Error("[CreateDepartmentHandler] error binding newDepartment data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	newDepartment.CompanyID = *companyID

	err = d.CreateDepartment(newDepartment)
	if err != nil {
		d.logger.Error("[CreateDepartmentHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error creating department",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "department created",
		Data:    nil,
	})
}

// updates department data, gets companyID from token, gets departmentID from param
func (d *Domain) UpdateDepartmentHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[UpdateDepartmentHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	departmentID, err := common.GetIDFromParam("departmentId", c)
	if err != nil {
		d.logger.Error("[UpdateDepartmentHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting department id from param",
			Data:    nil,
		})
	}

	dataToUpdate := new(schema.Department)

	err = c.Bind(dataToUpdate)
	if err != nil {
		d.logger.Error("[UpdateDepartmentHandler] error binding dataToUpdate", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	err = d.UpdateDepartment(companyID, departmentID, dataToUpdate)
	if err != nil {
		d.logger.Error("[UpdateDepartmentHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error updating department",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "department updated",
		Data:    nil,
	})
}

// deletes department data, gets companyID from token, gets departmentID from param
func (d *Domain) DeleteDepartmentHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[DeleteDepartmentHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	departmentID, err := common.GetIDFromParam("departmentId", c)
	if err != nil {
		d.logger.Error("[DeleteDepartmentHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting department id from param",
			Data:    nil,
		})
	}

	err = d.DeleteDepartment(companyID, departmentID)
	if err != nil {
		d.logger.Error("[DeleteDepartmentHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error deleting department",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "department has been deleted",
		Data:    nil,
	})
}

//! Location ------------------------------------------------------------

// creates location, gets companyID from token
func (d *Domain) CreateLocationHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[CreateLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}
	if companyID == nil {
		d.logger.Error("[CreateLocationHandler] companyID is nil")
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "company id is nil",
			Data:    nil,
		})
	}

	newLocation := new(schema.Location)

	err = c.Bind(newLocation)
	if err != nil {
		d.logger.Error("[CreateLocationHandler] error binding newLocation data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	newLocation.CompanyID = *companyID

	err = d.CreateLocation(newLocation)
	if err != nil {
		d.logger.Error("[CreateLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error creating location",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "location created",
		Data:    nil,
	})
}

// updates location data, gets companyID from token, gets locationID from param
func (d *Domain) UpdateLocationHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[UpdateLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	locationID, err := common.GetIDFromParam("locationId", c)
	if err != nil {
		d.logger.Error("[UpdateLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting location id from param",
			Data:    nil,
		})
	}

	dataToUpdate := new(schema.Location)

	err = c.Bind(dataToUpdate)
	if err != nil {
		d.logger.Error("[UpdateLocationHandler] error binding dataToUpdate", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	err = d.UpdateLocation(companyID, locationID, dataToUpdate)
	if err != nil {
		d.logger.Error("[UpdateLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error updating location",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "location data updated",
		Data:    nil,
	})
}

// allows manager to update location data, gets companyID and location ID from Param
func (d *Domain) ManagerUpdateLocationHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[ManagerUpdateLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	locationID, err := common.GetLocationIDFromToken(c)
	if err != nil {
		d.logger.Error("[ManagerUpdateLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting location id from token",
			Data:    nil,
		})
	}

	dataToUpdate := new(schema.Location)

	err = c.Bind(dataToUpdate)
	if err != nil {
		d.logger.Error("[ManagerUpdateLocationHandler] error binding dataToUpdate", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	err = d.UpdateLocationNoHeadOffice(companyID, locationID, dataToUpdate)
	if err != nil {
		d.logger.Error("[ManagerUpdateLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error updating location",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "location data updated",
		Data:    nil,
	})

}

// deletes location data, gets companyID from token, gets locationID from param
func (d *Domain) DeleteLocationHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[DeleteLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	locationID, err := common.GetIDFromParam("locationId", c)
	if err != nil {
		d.logger.Error("[DeleteLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting location id from param",
			Data:    nil,
		})
	}

	err = d.DeleteLocation(companyID, locationID)
	if err != nil {
		d.logger.Error("[DeleteLocationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error deleting location",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "location deleted",
		Data:    nil,
	})
}

//! Position ------------------------------------------------------------

// creates position, gets companyID from token
func (d *Domain) CreatePositionHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[CreatePositionHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}
	// to avoid nil pointer error
	if companyID == nil {
		d.logger.Error("[CreatePositionHandler] companyID is nil")
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "company id is nil",
			Data:    nil,
		})
	}

	newPosition := new(schema.Position)

	err = c.Bind(newPosition)
	if err != nil {
		d.logger.Error("[CreatePositionHandler] error binding newPosition data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	newPosition.CompanyID = *companyID

	err = d.CreatePosition(newPosition)
	if err != nil {
		d.logger.Error("[CreatePositionHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error creating position",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "position data has been created",
		Data:    nil,
	})
}

// updates position data, gets companyID from token, gets positionID from param
func (d *Domain) UpdatePositionHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[UpdatePositionHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	positionID, err := common.GetIDFromParam("positionId", c)
	if err != nil {
		d.logger.Error("[UpdatePositionHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting position id from param",
			Data:    nil,
		})
	}

	dataToUpdate := new(schema.Position)

	err = c.Bind(dataToUpdate)
	if err != nil {
		d.logger.Error("[UpdatePositionHandler] error binding dataToUpdate", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	err = d.UpdatePosition(companyID, positionID, dataToUpdate)
	if err != nil {
		d.logger.Error("[UpdatePositionHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error updating position",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "position updated",
		Data:    nil,
	})
}

// deletes position data, gets companyID from token, gets positionID from param
func (d *Domain) DeletePositionHandler(c echo.Context) error {
	companyID, err := common.GetCompanyIDFromToken(c)
	if err != nil {
		d.logger.Error("[DeletePositionHandlers]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting company id from token",
			Data:    nil,
		})
	}

	positionID, err := common.GetIDFromParam("positionId", c)
	if err != nil {
		d.logger.Error("[DeletePositionHandlers]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error getting position id from param",
			Data:    nil,
		})
	}

	err = d.DeletePosition(companyID, positionID)
	if err != nil {
		d.logger.Error("[DeletePositionHandlers]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "error deleting position",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "position deleted",
		Data:    nil,
	})
}
