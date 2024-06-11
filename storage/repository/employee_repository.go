package repository

import (
	"database/sql"
	"time"

	"github.com/karankap00r/employee_portal/storage/model"
)

//go:generate mockgen -source=employee_repository.go -destination=mocks/mock_employee_repository.go -package=mocks
type EmployeeRepository interface {
	GetAll() ([]model.Employee, error)
	GetByID(id int) (*model.Employee, error)
	Create(employee *model.Employee) error
	Update(id int, employee *model.Employee) error
}

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) GetAll() ([]model.Employee, error) {
	rows, err := r.db.Query("SELECT id, employee_id, name, position, email, salary, created_at, updated_at FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		if err := rows.Scan(&employee.ID, &employee.EmployeeID, &employee.Name, &employee.Position, &employee.Email, &employee.Salary, &employee.CreatedAt, &employee.UpdatedAt); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (r *employeeRepository) GetByID(id int) (*model.Employee, error) {
	var employee model.Employee
	err := r.db.QueryRow("SELECT id, employee_id, name, position, email, salary, created_at, updated_at FROM employees WHERE id = ?", id).
		Scan(&employee.ID, &employee.EmployeeID, &employee.Name, &employee.Position, &employee.Email, &employee.Salary, &employee.CreatedAt, &employee.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &employee, nil
}

func (r *employeeRepository) Create(employee *model.Employee) error {
	employee.CreatedAt = time.Now()
	employee.UpdatedAt = time.Now()
	_, err := r.db.Exec("INSERT INTO employees (employee_id, name, position, email, salary, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		employee.EmployeeID, employee.Name, employee.Position, employee.Email, employee.Salary, employee.CreatedAt, employee.UpdatedAt)
	return err
}

func (r *employeeRepository) Update(id int, employee *model.Employee) error {
	employee.UpdatedAt = time.Now()
	_, err := r.db.Exec("UPDATE employees SET name = ?, position = ?, email = ?, salary = ?, updated_at = ? WHERE id = ?",
		employee.Name, employee.Position, employee.Email, employee.Salary, employee.UpdatedAt, id)
	return err
}
