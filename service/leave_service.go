package service

import (
	"context"
	"fmt"
	"github.com/karankap00r/employee_portal/middleware"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/storage/repository"
	"time"
)

type LeaveService interface {
	GetLeaveBalance(ctx context.Context, employeeID string) (*model.LeaveBalance, error)
	RaiseLeaveRequest(context context.Context, request model.LeaveRequest) error
	UpdateLeave(ctx context.Context, id int, status, approvedBy string) error
	GetLeavesInRange(ctx context.Context, startDate, endDate time.Time) ([]model.LeaveRequest, error)
}

type leaveService struct {
	repo repository.LeaveRepository
}

func NewLeaveService(repo repository.LeaveRepository) LeaveService {
	return &leaveService{repo}
}

func (s *leaveService) GetLeaveBalance(ctx context.Context, employeeID string) (*model.LeaveBalance, error) {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.GetLeaveBalance(orgID, employeeID)
}

func (s *leaveService) RaiseLeaveRequest(ctx context.Context, request model.LeaveRequest) error {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("org ID not found in ctx")
	}

	request.OrgID = orgID
	return s.repo.CreateLeaveRequest(request)
}

func (s *leaveService) UpdateLeave(ctx context.Context, leaveID int, status, approvedBy string) error {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.UpdateLeaveRequestStatus(orgID, leaveID, status, approvedBy)
}

func (s *leaveService) GetLeavesInRange(ctx context.Context, startDate, endDate time.Time) ([]model.LeaveRequest, error) {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.GetLeavesInRange(orgID, startDate, endDate)
}
