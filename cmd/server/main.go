package main

import (
	serverconfig "GophKeeper/internal/config/server"
	"GophKeeper/internal/logger"
	"GophKeeper/internal/server/delivery/httpapi/handlers"
	"GophKeeper/internal/server/delivery/httpapi/middleware"
	"GophKeeper/internal/server/delivery/httpapi/router"
	"GophKeeper/internal/server/repository"
	"GophKeeper/internal/server/repository/database/postgresql"
	"GophKeeper/internal/server/usecase"
	"log"
	"net/http"
	"time"
)

func main() {
	var (
		keeperStorage repository.UserRepository
		err           error
	)

	// Запуск pprof на localhost:6060 для профилирования
	//go Pprof()

	log := logger.NewLogger()
	log.Info("starting server")

	cfg := serverconfig.NewConfigServerMock(log)

	// Попытка подключения к БД PostgreSQL несколько раз.
	for i := 0; i < 3; i++ {
		//Создание переменной для работы с БД
		keeperStorage, err = postgresql.NewPostgresStorage(cfg.FlagDatabaseDsn)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	uc := usecase.NewUseCase(keeperStorage, cfg, log)
	handle := handlers.NewHandlers(uc, log)
	mw := middleware.NewMiddleware(log)
	r := router.NewRouter(log, mw, handle)
	r.SetupRouter()

	err = http.ListenAndServe("localhost:6969", r.Mux)
	if err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
		return
	}
}

func Pprof() {
	log.Println("pprof запущен на :6060")
	if err := http.ListenAndServe("localhost:6060", nil); err != nil {
		log.Fatalf("Не удалось запустить pprof: %v", err)
	}
}
