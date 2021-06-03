package entities

type Projects struct {
	ID   int64   `json:"id"`   // идентификатор
	Lead *string `json:"lead"` // тимлид
	Name string  `json:"name"` // название проекта
	Desc string  `json:"desc"` // описание проекта
}
