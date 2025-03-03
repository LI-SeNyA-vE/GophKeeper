package router

import (
	"GophKeeper/internal/server/repository"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Router struct {
	log     *logrus.Entry
	storage repository.UserRepository
	*chi.Mux
}

func NewRouter(log *logrus.Entry, storage repository.UserRepository) *Router {
	return &Router{
		log:     log,
		storage: storage,
		Mux:     nil,
	}
}
