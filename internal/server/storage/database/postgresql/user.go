package postgresql

import "database/sql"

type DbUser struct {
	db *sql.DB
}

func (d *DbUser) RegistrationUser(login string, password string) error {
	d.db.Exec("INSERT INTO users (login, password) VALUES ($1, $2)", login, password)
	//TODO implement me
	panic("implement me")
}

func (d *DbUser) AuthorizationUser(login string, password string) error {

	//TODO implement me
	panic("implement me")
}
