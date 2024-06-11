package model

import "time"

type EmailReport struct {
	ID            int       `json:"id"`
	OrgID         int       `json:"org_id"`
	ReportType    string    `json:"report_type"`
	CronFrequency string    `json:"cron_frequency"`
	Status        string    `json:"status"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
