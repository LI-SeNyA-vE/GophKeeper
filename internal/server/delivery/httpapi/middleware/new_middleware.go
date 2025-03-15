package middleware

import "github.com/sirupsen/logrus"

// NewMiddleware создаёт новый объект Middleware с заданным логгером и конфигурацией сервера.
// Полученный объект затем может быть использован для инициализации различных middleware-функций,
// которые внедряются в цепочку обработки HTTP-запросов.
func NewMiddleware(log *logrus.Entry) *Middleware {
	return &Middleware{
		log: log,
	}
}
