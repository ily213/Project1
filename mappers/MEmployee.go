package mappers

import (
	"Project/app/models/entities"
	"database/sql"
	"time"
)

type MEmployee struct {
	db *sql.DB
}

// func NewMEmployee() *MEmployee{
// 	return &MEmployee {
// 		db:
// 	}
// }

type EmployeeDBType struct {
	Pk_id        int64          // идентификатор
	E_firstname  string         // имя
	E_lastname   string         // фамилия
	E_middlename sql.NullString // отчество
	E_Date       time.Time      // дата рождения
	E_adress     string         // адрес
}

// func (e EmployeeDBType) getSampleEmployee(){
// 	return new &Employee {
// 		ID: e.Pk_id
// 		E_midlename: e.
// 	}
// }

func (e *MEmployee) getEmployees() (employees []*entities.Employees, err error) {
	return
}

func (e *MEmployee) getEmployeeById(id int64) (employee *entities.Employees, err error) {
	return
}

func (e *MEmployee) createEmployee(employee_ *EmployeeDBType) (employee *entities.Employees, err error) {
	return
}

func (e *MEmployee) updateEmployee(employee_ *EmployeeDBType) (employee *entities.Employees, err error) {
	return
}

func (e *MEmployee) deleteEmployee(employee_ *EmployeeDBType) (err error) {
	return
}
