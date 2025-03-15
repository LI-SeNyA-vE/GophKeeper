package postgresql

const (
	queryTableUserIsNot = `
	CREATE TABLE IF NOT EXISTS "users" (
		id BIGSERIAL PRIMARY KEY,    
		uuid uuid UNIQUE NOT NULL,  
		login VARCHAR(255) UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`
	querySearchUser        = `SELECT id, uuid FROM users WHERE login = $1`
	queryRegistrationUser  = `INSERT INTO users(uuid, login, password) VALUES($1, $2, $3) RETURNING id, uuid`
	queryAuthorizationUser = `SELECT id, uuid, password FROM users WHERE login = $1`
	queryDeleteUser        = `DELETE FROM users WHERE uuid = $1`
)
