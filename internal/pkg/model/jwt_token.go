package model

import "github.com/dgrijalva/jwt-go"

type JwtToken struct {
	jwt.StandardClaims
	Type string `json:"type"`
}
