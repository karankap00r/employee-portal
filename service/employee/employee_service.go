package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/karankap00r/employee_portal/dto/request"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/storage/repository"
)

//go:generate mockgen -source=service/employee_service.go -destination=mocks/mock_employee_service.go -package=mocks
type EmployeeService interface {
	GetAllEmployees() ([]model.Employee, error)
	GetEmployeeByID(id int) (*model.Employee, error)
	CreateEmployee(request request.CreateEmployeeRequest) (*model.Employee, error)
	UpdateEmployeeByID(id int, request request.UpdateEmployeeByEmployeeIDRequest) (*model.Employee, error)
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo}
}

func (s *employeeService) GetAllEmployees() ([]model.Employee, error) {
	return s.repo.GetAll()
}

func (s *employeeService) GetEmployeeByID(id int) (*model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) CreateEmployee(request request.CreateEmployeeRequest) (*model.Employee, error) {
	employee := &model.Employee{
		EmployeeID: generateRandomEmployeeID(),
		Name:       request.Name,
		Position:   request.Position,
		Email:      request.Email,
		Salary:     request.Salary,
	}
	err := s.repo.Create(employee)
	return employee, err
}

func (s *employeeService) UpdateEmployeeByID(id int, request request.UpdateEmployeeByEmployeeIDRequest) (*model.Employee, error) {
	employee := &model.Employee{
		Name:     request.Name,
		Position: request.Position,
		Email:    request.Email,
		Salary:   request.Salary,
	}
	err := s.repo.Update(id, employee)
	if err != nil {
		return nil, err
	}
	employee.ID = id
	employee.EmployeeID = request.EmployeeID
	return employee, nil
}

func generateRandomEmployeeID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
