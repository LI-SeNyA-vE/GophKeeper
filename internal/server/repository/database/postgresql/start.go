package postgresql

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"

	"fmt"
	"log"
)

func NewPostgresStorage(dbConnect string) (*DataBase, error) {
	// Подключаемся к базе GopheKeeper
	db, err := sql.Open("pgx", dbConnect) //"postgresql://Senya:1q2w3e4r5t@localhost:5432/GopheKeeper?sslmode=disable"
	if err != nil {
		err = fmt.Errorf("ошибка подключения к системной базе данных: %v", err)
		log.Print(err)
		return nil, err
	}

	// Проверка соединения (ping)
	if err = db.Ping(); err != nil {
		err = fmt.Errorf("не удалось установить соединение с базой данных metrics: %w", err)
		log.Print(err)
		return nil, err
	}

	err = createTableIsNot(db)
	if err != nil {
		return nil, err
	}

	return &DataBase{
		db: db,
	}, nil
}
