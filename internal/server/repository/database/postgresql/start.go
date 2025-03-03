package postgresql

import (
	"database/sql"
	_ "database/sql"
	"log"
)

func NewPostgresStorage(dbConnect string) (*DataBase, error) {
	// Подключаемся к базе GopheKeeper
	db, err := sql.Open("pgx", dbConnect) //"postgresql://Senya@localhost:5432/GopheKeeper?sslmode=disable"
	if err != nil {
		log.Printf("ошибка подключения к системной базе данных: %v", err)
		return nil, err
	}

	// Проверка соединения (ping)
	if err = db.Ping(); err != nil {
		log.Printf("не удалось установить соединение с базой данных metrics: %v", err)
		return nil, err
	}

	return &DataBase{
		db: db,
	}, nil
}
