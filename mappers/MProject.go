package mappers

import (
	"Project/app/models/entities"
	"database/sql"
)

type MProject struct {
	db *sql.DB
}

type ProjectDBType struct {
	Pk_id        int64  // идентификатор
	Fk_lead      int64  // FK на тимлида
	Project_name string // название
	Project_desc string // описание
}

func (p *MProject) getProjects() (projects []*entities.Projects, err error) {
	return
}

func (p *MProject) getProjectById(id int64) (project *entities.Projects, err error) {
	return
}

func (p *MProject) createProject(project_ *ProjectDBType) (project *entities.Projects, err error) {
	return
}

func (p *MProject) updateProject(project_ *ProjectDBType) (project *entities.Projects, err error) {
	return
}

func (p *MProject) deleteProject(project_ *ProjectDBType) (err error) {
	return
}
