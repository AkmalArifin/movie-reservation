package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = os.Getenv("JWT_SECRET_KEY")

func GenerateToken(id int64, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString([]byte(key))
}

func GenerateRefreshToken(id int64, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   id,
		"role": role,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(key))
}

func VeryifyToken(token string) (int64, string, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(key), nil
	})

	if err != nil {
		return 0, "", errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return 0, "", errors.New("token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, "", errors.New("invalid token claims")
	}

	id := int64(claims["id"].(float64))
	role := claims["role"].(string)

	return id, role, nil

}
