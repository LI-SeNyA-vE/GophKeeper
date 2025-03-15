package usecase

import (
	"GophKeeper/internal/config/server"
	"GophKeeper/internal/server/domain"
	"GophKeeper/internal/server/repository"
	"GophKeeper/pkg/hashing"
	"GophKeeper/pkg/jwttoken"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type UserUseCase interface {
	// RegisterUser Принимает данные пользователя, регистрирует (добавление в базу) и возвращает access и refresh токены
	RegisterUser(user domain.User) (token *jwttoken.TokenDetails, err error)
	// AuthenticateUser Принимает данные пользователя и возвращает access и refresh токены
	AuthenticateUser(user domain.User) (token *jwttoken.TokenDetails, err error)
}

type userUseCase struct {
	repo      repository.UserRepository
	cfgServer *serverconfig.ConfigServer
	log       *logrus.Entry
}

func newUserUseCase(repo repository.UserRepository, cfgServer *serverconfig.ConfigServer, log *logrus.Entry) UserUseCase {
	return &userUseCase{
		repo:      repo,
		cfgServer: cfgServer,
		log:       log,
	}
}

func (u *userUseCase) RegisterUser(user domain.User) (token *jwttoken.TokenDetails, err error) {
	_, err = u.repo.SearchUser(user.Login)
	switch {
	case !errors.Is(err, sql.ErrNoRows):
		err = fmt.Errorf("пользователь с логином: %s уже существует. Ему не нужна регистрация", user.Login)
		u.log.Error(err)
		return nil, err
	case err != nil && !errors.Is(err, sql.ErrNoRows):
		err = fmt.Errorf("ошибка: %w при выполнении поиска пользователя по логину: %s", err, user.Login)
		u.log.Error(err)
		return nil, err
	}

	hashPassword, err := hashing.HashString(user.Password)
	if err != nil {
		err = fmt.Errorf("ошибка создание HASH пароля: %w", err)
		u.log.Error(err)
		return nil, err
	}

	newUser, err := u.repo.RegistrationUser(uuid.New().String(), user.Login, hashPassword)
	if err != nil {
		err = fmt.Errorf("ошибка при создании пользователя в БД: %v", err)
		u.log.Error(err)
		return nil, err
	}

	accessAndRefreshToken, err := jwttoken.NewToken(
		newUser.Uuid,
		u.cfgServer.FlagAccessKey,
		u.cfgServer.FlagRefreshKey,
		time.Minute*15,
		time.Hour*24,
	)
	if err != nil {
		err = fmt.Errorf("ошибка при создание access и refresh токенов: %w", err)
		u.log.Error(err)
		u.repo.DeleteUser(newUser.Uuid)
		u.log.Infof("Удалили пользователя '%v' из бд, в связи с ошибкой создания токенов", user.Login)
		return nil, err
	}

	return accessAndRefreshToken, nil
}

func (u *userUseCase) AuthenticateUser(user domain.User) (token *jwttoken.TokenDetails, err error) {
	userBD, err := u.repo.AuthorizationUser(user.Login)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		err = fmt.Errorf("пользователь с логином: %s ещё не существует. Ему нужна регистрация", user.Login)
		u.log.Error(err)
		return nil, err

	case err != nil && !errors.Is(err, sql.ErrNoRows):
		err = fmt.Errorf("ошибка: %w при получение данных из БД: %s", err, user.Login)
		u.log.Error(err)
		return nil, err
	}

	if !hashing.CheckString(userBD.Password, user.Password) {
		err = fmt.Errorf("пароли не совпадают")
		u.log.Error(err)
		return nil, err
	}

	accessAndRefreshToken, err := jwttoken.NewToken(
		userBD.Uuid,
		u.cfgServer.FlagAccessKey,
		u.cfgServer.FlagRefreshKey,
		time.Minute*15,
		time.Hour*24,
	)
	if err != nil {
		err = fmt.Errorf("ошибка при создание access и refresh токенов: %w", err)
		u.log.Error(err)
		return nil, err
	}

	return accessAndRefreshToken, nil
}
