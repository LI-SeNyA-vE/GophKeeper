package postgresql

import (
	"database/sql"
	"fmt"
)

func createTableIsNot(db *sql.DB) (err error) {
	query, err := db.Query(queryTableUserIsNot)
	if err != nil {
		return fmt.Errorf("ошибка при создание таблице юзера: %w", err)
	}
	query.Close()
	return nil
}
