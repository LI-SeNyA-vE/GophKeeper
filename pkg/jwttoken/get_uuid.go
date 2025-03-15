package jwttoken

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

func GetUuidUser(token *jwt.Token) (uuid string, err error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	uuid, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("uuid is empty")
	}
	return uuid, nil
}
