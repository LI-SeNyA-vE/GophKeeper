package jwttoken

import "errors"

var (
	LifetimeIsDead = errors.New("LifetimeIsDead")
	NoValid        = errors.New("NoValid")
)

// TokenDetails описывает структуру access и refresh токена
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
}
