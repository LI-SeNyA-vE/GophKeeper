// Package repository определяет общий интерфейс для работы с хранилищем паролей (UserRepository)
package repository

import "errors"

var (
	ErrNotFound = errors.New("по переданным данным, в базе нет записи")
)

type UserRepository interface {
	SearchUser(login string, password string) error
	RegistrationUser(login string, password string) error
	AuthorizationUser(login string, password string) error
}

type UserCred struct {
	Login    string
	Password string
}
