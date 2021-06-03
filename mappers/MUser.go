package mappers

import (
	"Project/app/models/entities"
	"database/sql"
)

type MUser struct {
	db *sql.DB
}

type UserDBType struct {
	Pk_id         int64  // идентификатор
	Fk_employeer  int64  // FK на сотрудника
	User_login    string // логин
	User_password string // пароль
}

func (u *MUser) login(user *UserDBType) (err error) {
	return
}

func (u *MUser) logout() (err error) {
	return
}

func (u *MUser) getCurrentEmployee() (employee *entities.Users, err error) {
	return
}
