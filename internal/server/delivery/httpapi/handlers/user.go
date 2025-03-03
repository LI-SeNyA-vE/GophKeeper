package handlers

import (
	"GophKeeper/internal/server/domain"
	"GophKeeper/internal/server/usecase"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserHandler interface {
	PostRegisterUser(w http.ResponseWriter, r *http.Request)
	PostAuthenticateUser(w http.ResponseWriter, r *http.Request)
}

// Реализация
type userHandler struct {
	uc  usecase.UserUseCase
	log *logrus.Entry
}

func NewUserHandler(uc usecase.UserUseCase, log *logrus.Entry) UserHandler {
	return &userHandler{
		uc:  uc,
		log: log,
	}
}

func (u userHandler) PostRegisterUser(w http.ResponseWriter, r *http.Request) {
	var userCred domain.User
	var buf bytes.Buffer
	var err error

	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		err = fmt.Errorf("ошибка чтения разбора тела запроса: %v", err)
		u.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(buf.Bytes(), &userCred)
	if err != nil {
		err = fmt.Errorf("ошибка разбора тела запроса в структуру: %v", err)
		u.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	accessToken, err := u.uc.RegisterUser(userCred)
	if err != nil {
		return
	}
	_ = accessToken

	//TODO implement me
	panic("implement me")
}

func (u userHandler) PostAuthenticateUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
