package common_models

import "github.com/dgrijalva/jwt-go"

// Claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
