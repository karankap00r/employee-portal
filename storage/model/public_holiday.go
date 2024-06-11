package model

type PublicHoliday struct {
	ID          int    `json:"id"`
	Country     string `json:"country"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Name        string `json:"name"`
	IsMandatory bool   `json:"is_mandatory"`
	CreatedBy   string `json:"created_by"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
