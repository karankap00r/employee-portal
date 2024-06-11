package service

import (
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/storage/repository"
)

type EmailReportService interface {
	GetActiveEmailReports(reportType string) ([]model.EmailReport, error)
}

type emailReportService struct {
	repo repository.EmailReportRepository
}

func NewEmailReportService(repo repository.EmailReportRepository) EmailReportService {
	return &emailReportService{repo}
}

func (s *emailReportService) GetActiveEmailReports(reportType string) ([]model.EmailReport, error) {
	return s.repo.GetActiveEmailReports(reportType)
}
