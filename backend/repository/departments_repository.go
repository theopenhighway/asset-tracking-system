package repository

import (
	"asset-tracking-system/entity"
	"database/sql"
)

type DeptRepostiory interface {
	CreateDepartment(name string) (*entity.Departments, error)
	GetDepartmentbyId(id int) (*entity.Departments, error)
	GetDepartment() (*[]entity.Departments, error)
	UpdateDepartment(id int, name string) (*entity.Departments, error)
	DeleteDepartment(id int) error
}

type deptRepository struct {
	db *sql.DB
}

func NewDeptRepository(db *sql.DB) *deptRepository {
	return &deptRepository{
		db: db,
	}
}

func (r *deptRepository) CreateDepartment(name string) (*entity.Departments, error) {
	registeredDept := entity.Departments{}

	sql := "INSERT INTO departments (name) VALUES ($1) RETURNING id, name"
	err := r.db.QueryRow(sql, name).Scan(&registeredDept.Id, &registeredDept.Name)

	if err != nil {
		return nil, err
	}

	return &registeredDept, nil
}

func (r *deptRepository) GetDepartmentbyId(id int) (*entity.Departments, error) {
	deptDetail := entity.Departments{}
	sql := "SELECT id, name from departments WHERE id = $1"
	err := r.db.QueryRow(sql, id).Scan(&deptDetail.Id, &deptDetail.Name)
	if err != nil {
		return nil, err
	}
	return &deptDetail, nil
}

func (r *deptRepository) GetDepartment() (*[]entity.Departments, error) {
	dept := []entity.Departments{}

	sql := "SELECT * FROM departments"
	rows, err := r.db.Query(sql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var deptRow entity.Departments

		err := rows.Scan(&deptRow.Id, &deptRow.Name)

		if err != nil {
			return &dept, err
		}

		dept = append(dept, deptRow)
	}

	err = rows.Err()

	if err != nil {
		return &dept, err
	}

	return &dept, nil
}

func (r *deptRepository) UpdateDepartment(id int, name string) (*entity.Departments, error) {
	dept := entity.Departments{}
	sql := "UPDATE departments SET name = $1 WHERE id = $2 RETURNING id, name"
	err := r.db.QueryRow(sql, id, name).Scan(&dept.Id, &dept.Name)

	if err != nil {
		return nil, err
	}

	return &dept, nil
}

func (r *deptRepository) DeleteDepartment(id int) error {
	sql := "DELETE departments WHERE id = $1"

	_, err := r.db.Query(sql, id)

	if err != nil {
		return err
	}

	return nil
}
