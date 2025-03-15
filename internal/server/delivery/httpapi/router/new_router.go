package router

import (
	"GophKeeper/internal/server/delivery/httpapi/handlers"
	"GophKeeper/internal/server/delivery/httpapi/middleware"
	"github.com/sirupsen/logrus"
)

func NewRouter(log *logrus.Entry, middleware *middleware.Middleware, handler *handlers.Handlers) *Router {
	return &Router{
		log:        log,
		middleware: middleware,
		handler:    handler,
		Mux:        nil,
	}
}
