package usecase

import (
	serverconfig "GophKeeper/internal/config/server"
	"GophKeeper/internal/server/repository"
	"github.com/sirupsen/logrus"
)

type UseCase struct {
	UserUseCase
}

func NewUseCase(repo repository.UserRepository, cfgServer *serverconfig.ConfigServer, log *logrus.Entry) UseCase {
	return UseCase{
		UserUseCase: newUserUseCase(repo, cfgServer, log),
	}
}
