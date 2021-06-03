package entities

import (
	"time"
)

type Employees struct {
	ID         int64     `json:"id"`         // идентификатор
	FirstName  string    `json:"firstname"`  // имя
	LastName   string    `json:"lastname"`   // фамилия
	MiddleName string    `json:"middlename"` // отчество
	Date       time.Time `json:"date"`       // дата рождения
	Adress     string    `json:"adress"`     // адрес
}
