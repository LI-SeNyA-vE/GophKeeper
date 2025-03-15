package handlers

import (
	"GophKeeper/internal/server/delivery/httpapi/handlers/userhandlers"
)

// Handlers хранит ссылки на логгер (logrus.Entry)
// и реализацию интерфейса UserRepository (repository).
// Используется в хендлерах для взаимодействия
// с базой, а также для логирования запросов/ответов.
type Handlers struct {
	userhandlers.UserHandler
}
