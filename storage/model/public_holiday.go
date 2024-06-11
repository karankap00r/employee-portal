package model

import "time"

type PublicHoliday struct {
	ID          int       `json:"id"`
	Country     string    `json:"country"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	IsMandatory bool      `json:"is_mandatory"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
