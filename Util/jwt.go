package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "usertoken"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 4).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("Could not parse token")
	}

	tokenIsValid := parseToken.Valid

	if !tokenIsValid {
		return errors.New("invalid token")
	}

	// claim, ok := parseToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("invalid token claim")
	// }

	// userEmail := claim["email"].(string)
	// userId := claim["userId"].(int64)

	return nil
}
