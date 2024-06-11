package service

import (
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/storage/repository"
)

// EmailReportService is an interface for the email report service
type EmailReportService interface {
	GetActiveEmailReports(reportType string) ([]model.EmailReport, error)
}

// emailReportService is a struct for the email report service
type emailReportService struct {
	repo repository.EmailReportRepository
}

// NewEmailReportService returns a new email report service
func NewEmailReportService(repo repository.EmailReportRepository) EmailReportService {
	return &emailReportService{repo}
}

// GetActiveEmailReports returns the active email reports
func (s *emailReportService) GetActiveEmailReports(reportType string) ([]model.EmailReport, error) {
	return s.repo.GetActiveEmailReports(reportType)
}
