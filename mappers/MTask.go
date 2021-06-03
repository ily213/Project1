package mappers

import (
	"Project/app/models/entities"
	"database/sql"
)

type MTask struct {
	db *sql.DB
}

type TaskDBType struct {
	Pk_id         int64         // идентификатор
	Fk_project    int64         // FK на проект
	Fk_type       int64         // FK на тип
	Fk_priority   int64         // FK на приоритет
	Fk_employeer  int64         // FK на сотрудника
	Fk_condition  int64         // FK на состояние
	Task_name     string        // название
	Task_desc     string        // описание
	Task_time     sql.NullInt64 // ожидаемое время выполнения
	Task_facttime sql.NullInt64 // фактическое время выполнения
}

func (t *MTask) getTasks() (tasks []*entities.Tasks, err error) {
	return
}

func (t *MTask) getTaskById(id int64) (task *entities.Tasks, err error) {
	return
}

func (t *MTask) createTask(task_ *TaskDBType) (task *entities.Tasks, err error) {
	return
}

func (t *MTask) updateTask(task_ *TaskDBType) (task *entities.Tasks, err error) {
	return
}

func (t *MTask) deleteTask(task_ *TaskDBType) (err error) {
	return
}
