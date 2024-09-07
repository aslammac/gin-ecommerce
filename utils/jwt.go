package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("your_secret_key") // Replace with a secure key in production

func GenerateJWT(phoneNumber string) (string, error) {
	claims := jwt.MapClaims{
		"phone_number": phoneNumber,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}