package service

import (
	"context"
	"fmt"
	"time"

	"github.com/karankap00r/employee_portal/middleware"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/storage/repository"
)

// RemoteWorkService is an interface for the remote work service
type RemoteWorkService interface {

	// GetRemoteWorkBalance returns the remote work balance for an employee
	GetRemoteWorkBalance(context context.Context, employeeID string) (*model.RemoteWorkBalance, error)

	// RaiseRemoteWorkRequest raises a remote work request
	RaiseRemoteWorkRequest(ctx context.Context, request model.RemoteWorkRequest) error

	// UpdateRemoteWorkRequestStatus updates a remote work request
	UpdateRemoteWorkRequestStatus(ctx context.Context, id int, status, approvedBy string) error

	// GetRemoteWorkRequestsInRange returns the remote work requests in a range
	GetRemoteWorkRequestsInRange(ctx context.Context, startDate, endDate time.Time) ([]model.RemoteWorkRequest, error)
}

// remoteWorkService is a struct for the remote work service
type remoteWorkService struct {
	repo repository.RemoteWorkRepository
}

// NewRemoteWorkService returns a new remote work service
func NewRemoteWorkService(repo repository.RemoteWorkRepository) RemoteWorkService {
	return &remoteWorkService{repo}
}

// GetRemoteWorkBalance returns the remote work balance for an employee
func (s *remoteWorkService) GetRemoteWorkBalance(ctx context.Context, employeeID string) (*model.RemoteWorkBalance, error) {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.GetRemoteWorkBalance(orgID, employeeID)
}

// RaiseRemoteWorkRequest raises a remote work request
func (s *remoteWorkService) RaiseRemoteWorkRequest(ctx context.Context, request model.RemoteWorkRequest) error {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("org ID not found in ctx")
	}

	request.OrgID = orgID
	return s.repo.CreateRemoteWorkRequest(request)
}

// UpdateRemoteWorkRequestStatus updates a remote work request
func (s *remoteWorkService) UpdateRemoteWorkRequestStatus(ctx context.Context, requestID int, status, updatedBy string) error {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.UpdateRemoteWorkRequestStatus(orgID, requestID, status, updatedBy)
}

// GetRemoteWorkRequestsInRange returns the remote work requests in a range
func (s *remoteWorkService) GetRemoteWorkRequestsInRange(ctx context.Context, startDate, endDate time.Time) ([]model.RemoteWorkRequest, error) {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.GetRemoteWorkRequestsInRange(orgID, startDate, endDate)
}
