package entities

type Users struct {
	ID       int64      `json:"id"`       // идентификатор
	Employee *Employees `json:"employee"` // сотрудник
	Login    string     `json:"login"`    // логин
	Password string     `json:"password"` // пароль
}
