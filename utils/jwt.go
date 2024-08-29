package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "secret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // type checking of value
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return errors.New("could not parse token")
	}
	if !parsedToken.Valid {
		return errors.New("invalid token")
	}
	//claims, ok := parsedToken.Claims.(jwt.MapClaims)
	//
	//if !ok {
	//	return errors.New("invalid token claims")
	//}
	//
	//userEmail := claims["email"].(string)
	//userId := claims["userId"].(int64)
	return nil
}
