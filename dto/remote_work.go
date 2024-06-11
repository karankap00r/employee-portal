package dto

// GetRemoteWorkBalanceRequest represents the request for getting remote work balance
type GetRemoteWorkBalanceRequest struct {
	OrgID      int    `json:"org_id"`
	EmployeeID string `json:"employee_id"`
}

// GetRemoteWorkBalanceResponse represents the response for getting remote work balance
type GetRemoteWorkBalanceResponse struct {
	EmployeeID    string `json:"employee_id"`
	Type          string `json:"type"`
	AnnualBalance int    `json:"annual_balance"`
}

// RaiseRemoteWorkRequest represents the request for raising remote work
type RaiseRemoteWorkRequest struct {
	OrgID      int    `json:"org_id"`
	EmployeeID string `json:"employee_id"`
	Type       string `json:"type"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Reason     string `json:"reason"`
}

// UpdateRemoteWorkRequest represents the request for updating remote work
type UpdateRemoteWorkRequest struct {
	OrgID     int    `json:"org_id"`
	RequestID int    `json:"request_id"`
	Status    string `json:"status"`
	UpdatedBy string `json:"updated_by"`
}

// GetRemoteWorkRequestsInRangeRequest represents the request for getting remote work requests in a range
type GetRemoteWorkRequestsInRangeRequest struct {
	OrgID     int    `json:"org_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

// GetRemoteWorkRequestsInRangeResponse represents the response for getting remote work requests in a range
type GetRemoteWorkRequestsInRangeResponse struct {
	ID         int    `json:"id"`
	EmployeeID string `json:"employee_id"`
	Type       string `json:"type"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Status     string `json:"status"`
}
