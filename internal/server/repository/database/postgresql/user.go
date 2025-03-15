package postgresql

import (
	"GophKeeper/internal/server/domain"
)

// SearchUser возвращает первое найденное совпадение по login (оно уникальное) в таблице users
func (d DataBase) SearchUser(login string) (userNoPass domain.User, err error) {
	err = d.db.QueryRow(querySearchUser, login).Scan(&userNoPass.Id, &userNoPass.Uuid)
	if err != nil {
		return userNoPass, err
	}
	return userNoPass, nil
}

func (d DataBase) RegistrationUser(uuid string, login string, password string) (newUser domain.User, err error) {
	err = d.db.QueryRow(queryRegistrationUser, uuid, login, password).Scan(&newUser.Id, &newUser.Uuid)
	if err != nil {
		return domain.User{}, err
	}
	newUser.Login = login
	return newUser, nil
}

func (d DataBase) AuthorizationUser(login string) (fullUser domain.User, err error) {
	err = d.db.QueryRow(queryAuthorizationUser, login).Scan(&fullUser.Id, &fullUser.Uuid, &fullUser.Password)
	if err != nil {
		return domain.User{}, err
	}
	return fullUser, nil
}

func (d DataBase) DeleteUser(uuid string) {
	d.db.QueryRow(queryDeleteUser, uuid)
}
