package repository

import (
	"database/sql"
	"log"
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
	var err error
	tx, err := r.db.Begin()
	defer func(err error) {
		if err != nil {
			if tx != nil {
				tx.Rollback()
			}
			log.Println(err)
		}
		tx.Commit()
	}(err)

	query := `INSERT INTO employees (org_id, employee_id, name, position, email, salary, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.Exec(query, employee.OrgID, employee.EmployeeID, employee.Name, employee.Position, employee.Email, employee.Salary, time.Now(), time.Now())
	if err != nil {
		return err
	}

	// Add logic to insert default leave balances based on org configurations
	leaveTypes := []string{"Sick Leave", "Vacation Leave"}
	leaveBalances := map[string]int{
		"Sick Leave":     10, // default value, can be replaced with org specific values
		"Vacation Leave": 15, // default value, can be replaced with org specific values
	}
	for _, leaveType := range leaveTypes {
		_, err := tx.Exec(`INSERT INTO leave_balances (org_id, employee_id, leave_type, annual_balance, created_at, updated_at)
		                  VALUES (?, ?, ?, ?, ?, ?)`,
			employee.OrgID, employee.EmployeeID, leaveType, leaveBalances[leaveType], time.Now(), time.Now())
		if err != nil {
			return err
		}
	}

	// Add logic to insert default remote work balances based on org configurations
	remoteWorkTypes := []string{"LOCAL", "CROSS_BORDER"}
	remoteWorkBalances := map[string]int{
		"LOCAL":        20, // default value, can be replaced with org specific values
		"CROSS_BORDER": 45, // default value, can be replaced with org specific values
	}
	for _, remoteWorkType := range remoteWorkTypes {
		_, err := tx.Exec(`INSERT INTO remote_work_balances (org_id, employee_id, type, annual_balance, created_at, updated_at)
		                  VALUES (?, ?, ?, ?, ?, ?)`,
			employee.OrgID, employee.EmployeeID, remoteWorkType, remoteWorkBalances[remoteWorkType], time.Now(), time.Now())
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *employeeRepository) Update(employeeID string, employee *model.Employee) error {
	employee.UpdatedAt = time.Now()
	_, err := r.db.Exec("UPDATE employees SET name = ?, position = ?, email = ?, salary = ?, updated_at = ? WHERE employee_id = ? and org_id = ?",
		employee.Name, employee.Position, employee.Email, employee.Salary, employee.UpdatedAt, employeeID, employee.OrgID)
	return err
}
