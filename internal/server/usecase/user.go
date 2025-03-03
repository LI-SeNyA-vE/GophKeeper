package usecase

import (
	"GophKeeper/internal/server/domain"
	"GophKeeper/internal/server/repository"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type UserUseCase interface {
	RegisterUser(user domain.User) (access string, err error)
	AuthenticateUser(user domain.User) (access string, err error)
}

type userUseCase struct {
	repo repository.UserRepository
	log  *logrus.Entry
}

func NewUserUseCase(repo repository.UserRepository, log *logrus.Entry) UserUseCase {
	return &userUseCase{
		repo: repo,
		log:  log,
	}
}

func (u *userUseCase) RegisterUser(user domain.User) (access string, err error) {
	err = u.repo.SearchUser(user.Login, user.Password)
	if !errors.Is(err, repository.ErrNotFound) {
		err = fmt.Errorf("данный пользователь уже существует %v", err)
		u.log.Error(err)
		return "", err
	}

	err = u.repo.RegistrationUser(user.Login, user.Password)
	if err != nil {
		err = fmt.Errorf("ошибка при создании пользователя в БД: %v", err)
		u.log.Error(err)
		return "", err
	}
	return
}

func (u *userUseCase) AuthenticateUser(user domain.User) (access string, err error) {
	//TODO implement me
	panic("implement me")
}
