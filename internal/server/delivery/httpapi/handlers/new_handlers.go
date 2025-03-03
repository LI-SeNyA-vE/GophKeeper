package handlers

import (
	"GophKeeper/internal/server/repository"
	"github.com/sirupsen/logrus"
)

// NewHandler создаёт новый Handler, инициализируя его логгером
// и интерфейсом хранилища метрик. Возвращает указатель на готовую структуру Handler,
// которую затем можно использовать в роутере (для регистрации HTTP-хендлеров).
func NewHandler(log *logrus.Entry, storage repository.UserRepository) *Handler {
	return &Handler{
		log:     log,
		storage: storage,
	}
}
