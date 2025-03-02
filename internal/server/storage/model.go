// Package storage определяет общий интерфейс для работы с хранилищем паролей (KeeperStorage)
package storage

type KeeperStorage interface {
	RegistrationUser(login string, password string) error
	AuthorizationUser(login string, password string) error
}
