package auth

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username  string `json:"username"`
	IsCreator bool   `json:"is_creator"`
	jwt.StandardClaims
}
