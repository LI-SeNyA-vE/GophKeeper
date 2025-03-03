package handlers

import (
	"GophKeeper/internal/server/storage"
	"github.com/sirupsen/logrus"
)

// Handler хранит ссылки на логгер (logrus.Entry)
// и реализацию интерфейса KeeperStorage (storage).
// Используется в хендлерах для взаимодействия
// с базой, а также для логирования запросов/ответов.
type Handler struct {
	log     *logrus.Entry
	storage storage.KeeperStorage
}

type Handlers interface {
}
