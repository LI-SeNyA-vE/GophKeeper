package router

import (
	"GophKeeper/internal/server/storage"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Router struct {
	log     *logrus.Entry
	storage storage.KeeperStorage
	*chi.Mux
}

func NewRouter(log *logrus.Entry, storage storage.KeeperStorage) *Router {
	return &Router{
		log:     log,
		storage: storage,
		Mux:     nil,
	}
}
