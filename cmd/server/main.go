package main

import (
	"GophKeeper/internal/logger"
	"GophKeeper/internal/server/storage"
	"GophKeeper/internal/server/storage/database/postgresql"
	"log"
	"net/http"
	"time"
)

func main() {
	var (
		keeperStorage storage.KeeperStorage
		err           error
	)

	// Запуск pprof на localhost:6060 для профилирования
	go Pprof()

	log := logger.NewLogger()
	log.Info("starting server")

	// Попытка подключения к БД PostgreSQL несколько раз.
	for i := 0; i < 3; i++ {
		//Создание переменной для работы с БД
		keeperStorage, err = postgresql.NewPostgresStorage("postgresql://Senya@localhost:5432/GopheKeeper?sslmode=disable")
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	_ = keeperStorage

	//Создание роутеров

}

func Pprof() {
	log.Println("pprof запущен на :6060")
	if err := http.ListenAndServe("localhost:6060", nil); err != nil {
		log.Fatalf("Не удалось запустить pprof: %v", err)
	}
}
