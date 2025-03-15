package router

import (
	"GophKeeper/internal/server/delivery/httpapi/handlers"
	"GophKeeper/internal/server/delivery/httpapi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Router struct {
	log        *logrus.Entry
	middleware *middleware.Middleware
	handler    *handlers.Handlers
	Mux        *chi.Mux
}
