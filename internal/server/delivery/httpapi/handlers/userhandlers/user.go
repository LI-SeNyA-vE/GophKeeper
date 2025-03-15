package userhandlers

import (
	"GophKeeper/internal/server/domain"
	"GophKeeper/internal/server/usecase"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type UserHandler interface {
	PostRegistrationUser(w http.ResponseWriter, r *http.Request)
	PostAuthorizationUser(w http.ResponseWriter, r *http.Request)
	PostLogoutUser(w http.ResponseWriter, r *http.Request)
	PostAddLoginAndPassword(w http.ResponseWriter, r *http.Request)
}

// Реализация
type userHandler struct {
	uc  usecase.UseCase
	log *logrus.Entry
}

// NewHandlers создаёт новый Handler, инициализируя его логгером
// и интерфейсом хранилища метрик. Возвращает указатель на готовую структуру Handler,
// которую затем можно использовать в роутере (для регистрации HTTP-хендлеров).
func NewHandlers(uc usecase.UseCase, log *logrus.Entry) UserHandler {
	return &userHandler{
		uc:  uc,
		log: log,
	}
}

func (u userHandler) PostRegistrationUser(w http.ResponseWriter, r *http.Request) {
	var userCred domain.User
	var buf bytes.Buffer
	var err error

	if r.Header.Get("Content-Type") != "application/json" {
		err = errors.New("Content-Type must be application/json")
		u.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		err = fmt.Errorf("ошибка чтения разбора тела запроса: %v", err)
		u.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &userCred)
	if err != nil {
		err = fmt.Errorf("ошибка разбора тела запроса в структуру: %v", err)
		u.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userCred.Login == "" || userCred.Password == "" {
		err = fmt.Errorf("логин или пароль не могут быть пустыми")
		u.log.Error("Был передан пользователь с пустым логином")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Переходим на следующий слой логики
	token, err := u.uc.RegisterUser(userCred)
	if err != nil {
		u.log.Errorf("кидаем ответ с ошибкой \"%v\"", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    token.RefreshToken,
		HttpOnly: true,                           // ❗ Доступен только серверу (защита от XSS)
		Secure:   false,                          // ❗ Только по HTTPS (обязательно в проде (true)) //TODO
		SameSite: http.SameSiteStrictMode,        // ❗ Защита от CSRF
		Path:     "/auth/refresh",                // ❗ Ограничиваем путь (только для refresh)
		Expires:  time.Now().Add(24 * time.Hour), // ❗ Refresh-токен живёт 7 дней
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Вы успешно зарегистрировались и вошли в систему"))
}

func (u userHandler) PostAuthorizationUser(w http.ResponseWriter, r *http.Request) {
	var userCred domain.User
	var buf bytes.Buffer
	var err error

	if r.Header.Get("Content-Type") != "application/json" {
		err = errors.New("Content-Type must be application/json")
		u.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = buf.ReadFrom(r.Body)
	if err != nil {
		err = fmt.Errorf("ошибка чтения разбора тела запроса: %v", err)
		u.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &userCred)
	if err != nil {
		err = fmt.Errorf("ошибка разбора тела запроса в структуру: %v", err)
		u.log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if userCred.Login == "" || userCred.Password == "" {
		err = fmt.Errorf("логин или пароль не могут быть пустыми")
		u.log.Error("Был передан пользователь с пустым логином")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := u.uc.AuthenticateUser(userCred)
	if err != nil {
		u.log.Errorf("кидаем ответ с ошибкой \"%v\"", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    token.RefreshToken,
		HttpOnly: true,                           // ❗ Доступен только серверу (защита от XSS)
		Secure:   false,                          // ❗ Только по HTTPS (обязательно в проде (true)) //TODO
		SameSite: http.SameSiteStrictMode,        // ❗ Защита от CSRF
		Path:     "/auth/refresh",                // ❗ Ограничиваем путь (только для refresh)
		Expires:  time.Now().Add(24 * time.Hour), // ❗ Refresh-токен живёт 7 дней
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Вы успешно авторизовались"))
}

func (u userHandler) PostLogoutUser(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/auth/refresh",
		MaxAge:   -1, // ❗ Удаляем куку
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Вы успешно вышли"))
}

func (u userHandler) PostAddLoginAndPassword(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
