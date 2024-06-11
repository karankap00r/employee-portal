package model

type OrgHoliday struct {
	ID        int    `json:"id"`
	OrgID     string `json:"org_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
}
