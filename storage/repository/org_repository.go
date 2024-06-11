package repository

import (
	"database/sql"
	"github.com/karankap00r/employee_portal/storage/model"
)

type OrgRepository interface {
	GetOrgByClientID(clientID string) (*model.Org, error)
}

type orgRepository struct {
	db *sql.DB
}

func NewOrgRepository(db *sql.DB) OrgRepository {
	return &orgRepository{db}
}

func (r *orgRepository) GetOrgByClientID(clientID string) (*model.Org, error) {
	query := `SELECT id, name, client_id, contact_email, contact_phone, created_at, updated_at FROM orgs WHERE client_id = ?`
	row := r.db.QueryRow(query, clientID)
	var org model.Org
	err := row.Scan(&org.ID, &org.Name, &org.ClientID, &org.ContactEmail, &org.ContactPhone, &org.CreatedAt, &org.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &org, nil
}
