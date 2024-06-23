package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/alsey89/people-matter/internal/common"
)

func (d *Domain) AuthenticateUserHandler(c echo.Context) error {
	creds := new(SigninCredentials)
	err := c.Bind(creds)
	if err != nil {
		d.logger.Error("[AuthenticateUserHandler] error binding credentials", zap.Error(err))
		return c.JSON(http.StatusBadRequest, common.APIResponse{
			Message: "invalid form data",
			Data:    nil,
		})
	}

	email := creds.Email
	password := creds.Password

	userRoles, err := d.AuthenticateUserService(email, password)
	switch {
	case err != nil:
		d.logger.Error("[AuthenticateUserHandler]", zap.Error(err))
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, common.APIResponse{
				Message: "user not found",
				Data:    nil,
			})
		case errors.Is(err, ErrUserNotConfirmed):
			return c.JSON(http.StatusForbidden, common.APIResponse{
				Message: "user not confirmed",
				Data:    nil,
			})
		case errors.Is(err, ErrInvalidCredentials):
			return c.JSON(http.StatusUnauthorized, common.APIResponse{
				Message: "invalid credentials",
				Data:    nil,
			})
		default:
			return c.JSON(http.StatusInternalServerError, common.APIResponse{
				Message: "something went wrong",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "user authenticated",
		Data:    userRoles,
	})
}

func (d *Domain) GenerateTokenHandler(c echo.Context) error {
	creds := new(SigninCredentials)
	err := c.Bind(creds)
	if err != nil {
		d.logger.Error("[GenerateTokenHandler] error binding credentials", zap.Error(err))
		return c.JSON(http.StatusBadRequest, common.APIResponse{
			Message: "invalid form data",
			Data:    nil,
		})
	}

	userRoleID := creds.UserRoleID

	preloadedUserRole, err := d.GetUserBySelectedUserRole(userRoleID)
	if err != nil {
		d.logger.Error("[GenerateTokenHandler] error fetching user", zap.Error(err))
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return c.JSON(http.StatusNotFound, common.APIResponse{
				Message: "user not found",
				Data:    nil,
			})
		default:
			return c.JSON(http.StatusInternalServerError, common.APIResponse{
				Message: "something went wrong",
				Data:    nil,
			})
		}
	}

	claims := jwt.MapClaims{
		"id":         preloadedUserRole.User.ID,
		"email":      preloadedUserRole.User.Email,
		"companyId":  preloadedUserRole.Company.ID,
		"role":       preloadedUserRole.Role.Name,
		"locationId": preloadedUserRole.LocationID,
		"userRoleId": preloadedUserRole.ID,
	}

	t, err := d.params.JWT.GenerateToken("jwt_auth", claims)
	if err != nil {
		d.logger.Error("[GenerateTokenHandler] error generating token", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}
	if t == nil {
		d.logger.Error("[GenerateTokenHandler] token is nil")
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = *t
	cookie.HttpOnly = true
	cookie.Secure = viper.GetBool("IS_PRODUCTION")
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * 72)

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "user has been signed in",
		Data:    preloadedUserRole,
	})
}

func (d *Domain) SignoutHandler(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = ""
	cookie.HttpOnly = true
	cookie.Secure = viper.GetBool("IS_PRODUCTION")
	cookie.Path = "/"
	cookie.Expires = time.Unix(0, 0) //* set the cookie to expire immediately

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "user has been signed out",
		Data:    nil,
	})
}

func (d *Domain) ConfirmationHandler(c echo.Context) error {
	// Assuming token has been validated by middleware and user set in context
	user, ok := c.Get("user").(*jwt.Token)
	// unexpected error
	if !ok || user == nil {
		return c.JSON(http.StatusUnauthorized, common.APIResponse{
			Message: "something went wrong with token validation",
			Data:    nil,
		})
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, common.APIResponse{
			Message: "error asserting claims",
			Data:    nil,
		})
	}

	floatID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(http.StatusUnauthorized, common.APIResponse{
			Message: "error asserting id",
			Data:    nil,
		})
	}

	uintID := uint(floatID)

	err := d.ConfirmEmailService(uintID)
	if err != nil {
		d.logger.Error("[ConfirmationHandler]", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, common.APIResponse{
			Message: "something went wrong",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "user has been confirmed",
		Data:    nil,
	})
}

func (d *Domain) CheckAuth(c echo.Context) error {
	// Assuming token has been validated by middleware and user set in context
	user, ok := c.Get("user").(*jwt.Token)
	// unexpected error
	if !ok || user == nil {
		return c.JSON(http.StatusUnauthorized, common.APIResponse{
			Message: "something went wrong with token validation",
			Data:    nil,
		})
	}

	_, ok = user.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, common.APIResponse{
			Message: "error asserting claims",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "authenticated",
		Data:    nil,
	})
}

func (d *Domain) GetCSRFToken(c echo.Context) error {
	return c.JSON(http.StatusOK, common.APIResponse{
		Message: "csrf token has been set",
		Data:    nil,
	})
}
