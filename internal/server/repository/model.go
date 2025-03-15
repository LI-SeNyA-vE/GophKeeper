// Package repository определяет общий интерфейс для работы с хранилищем паролей (UserRepository)
package repository

import (
	"GophKeeper/internal/server/domain"
)

//var (
//	ErrNotFound = errors.New("по переданным данным, в базе нет записи")
//)

type UserRepository interface {
	SearchUser(login string) (userNoPass domain.User, err error)
	RegistrationUser(uuid string, login string, password string) (newUser domain.User, err error)
	AuthorizationUser(login string) (fullUser domain.User, err error)
	DeleteUser(uuid string)
}
