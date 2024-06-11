package model

type Org struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ClientID     string `json:"client_id"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
