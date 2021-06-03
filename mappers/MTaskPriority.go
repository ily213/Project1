package mappers

import (
	"Project/app/models/entities"
	"database/sql"
)

type MTaskPriority struct {
	db *sql.DB
}

type MTaskPriorityDBType struct {
	Pk_id         int64  // идентификатор
	priority_name string // название
}

func (tc *MTaskPriority) getPriorityByID(id int64) (tasks []*entities.TaskPriority, err error) {
	return
}

func (tc *MTaskPriority) getIDByPriority() (tasks []*entities.TaskPriority, err error) {
	return
}
