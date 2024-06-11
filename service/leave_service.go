package service

import (
	"context"
	"fmt"
	"github.com/karankap00r/employee_portal/middleware"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/storage/repository"
	"time"
)

// LeaveService is an interface for the leave service
type LeaveService interface {

	// GetLeaveBalance returns the leave balance for an employee
	GetLeaveBalance(ctx context.Context, employeeID string) (*model.LeaveBalance, error)

	// RaiseLeaveRequest raises a leave request
	RaiseLeaveRequest(context context.Context, request model.LeaveRequest) error

	// UpdateLeave updates a leave request
	UpdateLeave(ctx context.Context, id int, status, approvedBy string) error

	// GetLeavesInRange returns the leaves in a range
	GetLeavesInRange(ctx context.Context, startDate, endDate time.Time) ([]model.LeaveRequest, error)
}

// leaveService is a struct for the leave service
type leaveService struct {
	repo repository.LeaveRepository
}

// NewLeaveService returns a new leave service
func NewLeaveService(repo repository.LeaveRepository) LeaveService {
	return &leaveService{repo}
}

// GetLeaveBalance returns the leave balance for an employee
func (s *leaveService) GetLeaveBalance(ctx context.Context, employeeID string) (*model.LeaveBalance, error) {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.GetLeaveBalance(orgID, employeeID)
}

// RaiseLeaveRequest raises a leave request
func (s *leaveService) RaiseLeaveRequest(ctx context.Context, request model.LeaveRequest) error {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("org ID not found in ctx")
	}

	request.OrgID = orgID
	return s.repo.CreateLeaveRequest(request)
}

// UpdateLeave updates a leave request
func (s *leaveService) UpdateLeave(ctx context.Context, leaveID int, status, approvedBy string) error {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.UpdateLeaveRequestStatus(orgID, leaveID, status, approvedBy)
}

// GetLeavesInRange returns the leaves in a range
func (s *leaveService) GetLeavesInRange(ctx context.Context, startDate, endDate time.Time) ([]model.LeaveRequest, error) {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.GetLeavesInRange(orgID, startDate, endDate)
}
