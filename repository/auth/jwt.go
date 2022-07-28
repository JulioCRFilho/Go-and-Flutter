package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("internalsecretinenviroment")

type JWT struct {
	Email string `json:"email" binding:"required"`
	Pass  string `json:"pass" binding:"required"`
	jwt.StandardClaims
}

func GenerateToken(email string, pass string) (string, error) {
	expiration := time.Now().Add(1 * time.Hour)
	claims := &JWT{
		email,
		pass,
		jwt.StandardClaims{
			ExpiresAt: expiration.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) error {
	if token, err := jwt.ParseWithClaims(
		signedToken,
		&JWT{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	); err != nil {
		return err
	} else {
		if claims, ok := token.Claims.(*JWT); !ok {
			return errors.New("couldn't parse claims")
		} else {
			if claims.ExpiresAt < time.Now().Local().Unix() {
				return errors.New("token expired")
			} else {
				return nil
			}
		}
	}
}
