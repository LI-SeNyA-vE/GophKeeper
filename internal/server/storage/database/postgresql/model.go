// Package postgresql реализует интерфейс KeeperStorage на основе базы данных PostgreSQL
package postgresql

import (
	"database/sql"
)

/*
DataBase представляет собой структуру, хранящую:

  - log: логгер на базе logrus.Entry,
  - db: объект *sql.DB (активное соединение с базой данных PostgreSQL).

Она реализует интерфейс KeeperStorage.
*/
type DataBase struct {
	db *sql.DB
}
