package middleware

import "github.com/sirupsen/logrus"

// Middleware содержит ссылку на общий логгер (logrus.Entry) и структуру конфигурации сервера.
// Используется при инициализации набора промежуточных обработчиков (LoggingMiddleware,
// HashSHA256, GunzipMiddleware и др.).
type Middleware struct {
	log *logrus.Entry
}
