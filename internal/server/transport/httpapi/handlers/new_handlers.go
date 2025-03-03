package handlers

import (
	"GophKeeper/internal/server/storage"
	"github.com/sirupsen/logrus"
)

// NewHandler создаёт новый Handler, инициализируя его логгером
// и интерфейсом хранилища метрик. Возвращает указатель на готовую структуру Handler,
// которую затем можно использовать в роутере (для регистрации HTTP-хендлеров).
func NewHandler(log *logrus.Entry, storage storage.KeeperStorage) *Handler {
	return &Handler{
		log:     log,
		storage: storage,
	}
}
