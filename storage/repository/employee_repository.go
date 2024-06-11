package repository

import (
	"database/sql"
	"time"

	"github.com/karankap00r/employee_portal/storage/model"
)

//go:generate mockgen -source=employee_repository.go -destination=mocks/mock_employee_repository.go -package=mocks
type EmployeeRepository interface {
	GetAll(orgID int) ([]model.Employee, error)
	GetByEmployeeID(orgID int, employeeID string) (*model.Employee, error)
	Create(employee *model.Employee) error
	Update(employeeID string, employee *model.Employee) error
}

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (r *employeeRepository) GetAll(orgID int) ([]model.Employee, error) {
	rows, err := r.db.Query("SELECT id, employee_id, name, position, email, salary, created_at, updated_at FROM employees where org_id = ?", orgID)
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

func (r *employeeRepository) GetByEmployeeID(orgID int, employeeID string) (*model.Employee, error) {
	var employee model.Employee
	err := r.db.QueryRow("SELECT id, employee_id, name, position, email, salary, created_at, updated_at FROM employees WHERE employee_id = ? and org_id = ?", employeeID, orgID).
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
	_, err := r.db.Exec("INSERT INTO employees (org_id, employee_id, name, position, email, salary, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		employee.OrgID, employee.EmployeeID, employee.Name, employee.Position, employee.Email, employee.Salary, employee.CreatedAt, employee.UpdatedAt)
	return err
}

func (r *employeeRepository) Update(employeeID string, employee *model.Employee) error {
	employee.UpdatedAt = time.Now()
	_, err := r.db.Exec("UPDATE employees SET name = ?, position = ?, email = ?, salary = ?, updated_at = ? WHERE employee_id = ? and org_id = ?",
		employee.Name, employee.Position, employee.Email, employee.Salary, employee.UpdatedAt, employeeID, employee.OrgID)
	return err
}
