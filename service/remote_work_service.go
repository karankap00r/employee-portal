package service

import (
	"context"
	"fmt"
	"time"

	"github.com/karankap00r/employee_portal/middleware"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/storage/repository"
)

type RemoteWorkService interface {
	GetRemoteWorkBalance(context context.Context, employeeID string) (*model.RemoteWorkBalance, error)
	RaiseRemoteWorkRequest(ctx context.Context, request model.RemoteWorkRequest) error
	UpdateRemoteWorkRequestStatus(ctx context.Context, id int, status, approvedBy string) error
	GetRemoteWorkRequestsInRange(ctx context.Context, startDate, endDate time.Time) ([]model.RemoteWorkRequest, error)
}

type remoteWorkService struct {
	repo repository.RemoteWorkRepository
}

func NewRemoteWorkService(repo repository.RemoteWorkRepository) RemoteWorkService {
	return &remoteWorkService{repo}
}

func (s *remoteWorkService) GetRemoteWorkBalance(ctx context.Context, employeeID string) (*model.RemoteWorkBalance, error) {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.GetRemoteWorkBalance(orgID, employeeID)
}

func (s *remoteWorkService) RaiseRemoteWorkRequest(ctx context.Context, request model.RemoteWorkRequest) error {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("org ID not found in ctx")
	}

	request.OrgID = orgID
	return s.repo.CreateRemoteWorkRequest(request)
}

func (s *remoteWorkService) UpdateRemoteWorkRequestStatus(ctx context.Context, requestID int, status, updatedBy string) error {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.UpdateRemoteWorkRequestStatus(orgID, requestID, status, updatedBy)
}

func (s *remoteWorkService) GetRemoteWorkRequestsInRange(ctx context.Context, startDate, endDate time.Time) ([]model.RemoteWorkRequest, error) {
	orgID, ok := middleware.GetOrgIDFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("org ID not found in ctx")
	}

	return s.repo.GetRemoteWorkRequestsInRange(orgID, startDate, endDate)
}
