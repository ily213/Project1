package mappers

import (
	"Project/app/models/entities"
	"database/sql"
)

type MTaskType struct {
	db *sql.DB
}

type MTaskTypeDBType struct {
	Pk_id     int64  // идентификатор
	Type_name string // название
}

func (tc *MTaskType) getTypeByID(id int64) (tasks []*entities.TaskType, err error) {
	return
}

func (tc *MTaskType) getIDByType() (tasks []*entities.TaskType, err error) {
	return
}
