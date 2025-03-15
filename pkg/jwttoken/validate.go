package jwttoken

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// ValidateToken Функция проверки токена
func ValidateToken(tokenString, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм подписи - HMAC-SHA256
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("ожидался другой метод подписи")
		}
		return []byte(secret), nil
	})

	// Проверяем сам токен на валидность
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("токен неволиден, возможно отредактирован/подменён: %w", NoValid)
	}
	if time.Now().Unix() > token.Claims.(jwt.MapClaims)["exp"].(int64) {
		return nil, fmt.Errorf("токен не действителен: %w", LifetimeIsDead)
	}

	return token, nil
}
