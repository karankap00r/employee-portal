package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/karankap00r/employee_portal/dto/request"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/storage/repository"
	"github.com/karankap00r/employee_portal/util"
)

const remoteTimezone = "America/New_York" // Specify the remote server timezone

//go:generate mockgen -source=service/employee_service.go -destination=mocks/mock_employee_service.go -package=mocks
type EmployeeService interface {
	GetAllEmployees() ([]model.Employee, error)
	GetEmployeeByEmployeeID(employeeID string) (*model.Employee, error)
	CreateEmployee(request request.CreateEmployeeRequest) (*model.Employee, error)
	UpdateEmployeeByEmployeeID(employeeID string, request request.UpdateEmployeeByEmployeeIDRequest) (*model.Employee, error)
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo}
}

func (s *employeeService) GetAllEmployees() ([]model.Employee, error) {
	employees, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	for _, employee := range employees {
		s.transformEmployeeToLocalTimezone(&employee)
	}
	return employees, nil
}

func (s *employeeService) GetEmployeeByEmployeeID(employeeID string) (*model.Employee, error) {
	employee, err := s.repo.GetByEmployeeID(employeeID)
	if err != nil {
		return nil, err
	}
	if employee != nil {
		s.transformEmployeeToLocalTimezone(employee)
	}
	return employee, nil
}

func (s *employeeService) CreateEmployee(request request.CreateEmployeeRequest) (*model.Employee, error) {
	currentTime, err := util.GetCurrentTimeInTimezone(remoteTimezone)
	if err != nil {
		return nil, err
	}

	employee := &model.Employee{
		EmployeeID: generateRandomEmployeeID(),
		Name:       request.Name,
		Position:   request.Position,
		Email:      request.Email,
		Salary:     request.Salary,
		CreatedAt:  currentTime,
		UpdatedAt:  currentTime,
	}
	err = s.repo.Create(employee)
	return employee, err
}

func (s *employeeService) UpdateEmployeeByEmployeeID(employeeID string, request request.UpdateEmployeeByEmployeeIDRequest) (*model.Employee, error) {
	currentTime, err := util.GetCurrentTimeInTimezone(remoteTimezone)
	if err != nil {
		return nil, err
	}

	employee := &model.Employee{
		Name:      request.Name,
		Position:  request.Position,
		Email:     request.Email,
		Salary:    request.Salary,
		UpdatedAt: currentTime,
	}

	err = s.repo.Update(employeeID, employee)
	if err != nil {
		return nil, err
	}
	employee.EmployeeID = request.EmployeeID
	return employee, nil
}

/* Helper Functions */

func generateRandomEmployeeID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func (s *employeeService) transformEmployeeToLocalTimezone(employee *model.Employee) error {
	localTimezone, err := util.GetLocalTimezone()
	if err != nil {
		return err
	}
	employee.CreatedAt, err = util.ConvertToTimezone(employee.CreatedAt, localTimezone)
	if err != nil {
		return err
	}
	employee.UpdatedAt, err = util.ConvertToTimezone(employee.UpdatedAt, localTimezone)
	if err != nil {
		return err
	}
	return nil
}
