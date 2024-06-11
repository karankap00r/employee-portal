package repository

import (
	"database/sql"

	"github.com/karankap00r/employee_portal/storage/model"
)

type EmailReportRepository interface {
	GetActiveEmailReports(reportType string) ([]model.EmailReport, error)
}

type emailReportRepository struct {
	db *sql.DB
}

func NewEmailReportRepository(db *sql.DB) EmailReportRepository {
	return &emailReportRepository{db}
}

func (r *emailReportRepository) GetActiveEmailReports(reportType string) ([]model.EmailReport, error) {
	query := `SELECT id, org_id, report_type, cron_frequency, status, email, created_at, updated_at
	          FROM email_reports
	          WHERE report_type = ? AND status = 'ACTIVE'`
	rows, err := r.db.Query(query, reportType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []model.EmailReport
	for rows.Next() {
		var report model.EmailReport
		err := rows.Scan(&report.ID, &report.OrgID, &report.ReportType, &report.CronFrequency, &report.Status, &report.Email, &report.CreatedAt, &report.UpdatedAt)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}
