package entities

type Tasks struct {
	ID        int64         `json:"id"`        // идентификатор
	Type      TaskType      `json:"type"`      // тип задачи
	Priority  TaskPriority  `json:"priority"`  // приоритет задачи
	Employeer *Employees    `json:"employeer"` // сотрудник
	Condition TaskCondition `json:"condition"` // состояние задачи
	Name      string        `json:"name"`      // название
	Desc      string        `json:"desc"`      // описание
	Time      int           `json:"time"`      // ожидаемое время выполнения
	FactTime  int           `json:"facttime"`  // фактическое время выполнения
}

type TaskType string      // Тип задачи
type TaskPriority string  // Приоритет задачи
type TaskCondition string // Состояние задачи

const (
	//Типы задач
	TASK_TYPE_TEST   TaskType = "Тест"
	TASK_TYPE_CREATE TaskType = "Создание"
	TASK_TYPE_FIX    TaskType = "Исправление ошибок"
	TASK_TYPE_CHANGE TaskType = "Изменение"

	//Приоритеты задач
	TASK_PRIORITY_LOW       TaskPriority = "Низкий"
	TASK_PRIORITY_MEDIUM    TaskPriority = "Средний"
	TASK_PRIORITY_HIGH      TaskPriority = "Высокий"
	TASK_PRIORITY_IMPORTANT TaskPriority = "Важно"

	//Состояния задач
	TASK_CONDITION_NEW      TaskCondition = "Новая"
	TASK_CONDITION_GIVEN    TaskCondition = "Назначена"
	TASK_CONDITION_INWORK   TaskCondition = "В работе"
	TASK_CONDITION_PAUSE    TaskCondition = "Пауза"
	TASK_CONDITION_COMPLETE TaskCondition = "Решено"
	TASK_CONDITION_TOCHANGE TaskCondition = "Согласовано"
)
