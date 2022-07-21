package auth

import "github.com/dgrijalva/jwt-go"

var jwtKey = []byte("internalsecretinenviroment")

type JWT struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
