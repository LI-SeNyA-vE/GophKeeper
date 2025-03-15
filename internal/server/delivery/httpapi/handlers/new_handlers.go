package handlers

import (
	"GophKeeper/internal/server/delivery/httpapi/handlers/userhandlers"
	"GophKeeper/internal/server/usecase"
	"github.com/sirupsen/logrus"
)

func NewHandlers(uc usecase.UseCase, log *logrus.Entry) *Handlers {
	return &Handlers{
		UserHandler: userhandlers.NewHandlers(uc, log),
	}
}
