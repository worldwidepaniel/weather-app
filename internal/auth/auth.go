package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var signingKey = []byte("thereisnocowlevel")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		err := fmt.Errorf("problem with generating token: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
