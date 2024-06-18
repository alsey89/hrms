package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

// * JWT claims struct
type Claims struct {
	ID         uint   `json:"id" bson:"id"`
	CompanyID  uint   `json:"companyId"`
	Role       string `json:"role" default:"user"`
	LocationID *uint  `json:"locationId"`
	Email      string `json:"email"`
	jwt.RegisteredClaims
}

type SigninCredentials struct {
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
	CompanyID uint   `json:"companyId" bson:"companyId"`
}

//------------------------------------------------------------
