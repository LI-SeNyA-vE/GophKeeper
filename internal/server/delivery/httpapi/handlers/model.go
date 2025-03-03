package handlers

import (
	"GophKeeper/internal/server/repository"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Handler хранит ссылки на логгер (logrus.Entry)
// и реализацию интерфейса UserRepository (repository).
// Используется в хендлерах для взаимодействия
// с базой, а также для логирования запросов/ответов.
type Handler struct {
	log     *logrus.Entry
	storage repository.UserRepository
}

type Handlers interface {
	PostRegistrationUser(w http.ResponseWriter, r *http.Request)
	PostAuthorizationUser(w http.ResponseWriter, r *http.Request)
}
