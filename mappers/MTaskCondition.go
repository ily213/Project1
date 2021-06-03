package mappers

import (
	"Project/app/models/entities"
	"database/sql"
)

type MTaskCondition struct {
	db *sql.DB
}

type MTaskConditionDBType struct {
	Pk_id          int64  // идентификатор
	Condition_name string // название
}

func (tc *MTaskCondition) getConditionByID(id int64) (taskConditions []*entities.TaskCondition, err error) {
	return
}

func (tc *MTaskCondition) getIDByCondition() (taskConditions []*entities.TaskCondition, err error) {
	return
}
