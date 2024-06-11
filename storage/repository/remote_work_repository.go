package repository

import (
	"database/sql"
	"github.com/karankap00r/employee_portal/common"
	"github.com/karankap00r/employee_portal/storage/model"
	"log"
	"time"
)

type RemoteWorkRepository interface {
	GetRemoteWorkBalance(orgID int, employeeID string) (*model.RemoteWorkBalance, error)
	CreateRemoteWorkRequest(request model.RemoteWorkRequest) error
	UpdateRemoteWorkRequestStatus(orgID, requestID int, status, approvedBy string) error
	GetRemoteWorkRequestsInRange(orgID int, startDate, endDate time.Time) ([]model.RemoteWorkRequest, error)
}

type remoteWorkRepository struct {
	db *sql.DB
}

func NewRemoteWorkRepository(db *sql.DB) RemoteWorkRepository {
	return &remoteWorkRepository{db}
}

func (r *remoteWorkRepository) GetRemoteWorkBalance(orgID int, employeeID string) (*model.RemoteWorkBalance, error) {
	query := `SELECT id, org_id, employee_id, type, annual_balance, created_at, updated_at FROM remote_work_balances WHERE employee_id = ? and org_id = ?`
	row := r.db.QueryRow(query, employeeID, orgID)
	var balance model.RemoteWorkBalance
	err := row.Scan(&balance.ID, &balance.OrgID, &balance.EmployeeID, &balance.Type, &balance.AnnualBalance, &balance.CreatedAt, &balance.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (r *remoteWorkRepository) CreateRemoteWorkRequest(request model.RemoteWorkRequest) error {
	query := `INSERT INTO remote_work_requests (org_id, employee_id, type, start_date, end_date, reason, status, approved_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, request.OrgID, request.EmployeeID, request.Type, request.StartDate, request.EndDate, request.Reason, request.Status, request.UpdatedBy, request.CreatedAt, request.UpdatedAt)
	return err
}

func (r *remoteWorkRepository) UpdateRemoteWorkRequestStatus(orgID, requestID int, status, updatedBy string) error {
	statusUpdateQuery := `UPDATE remote_work_requests SET status = ?, updated_by = ?, updated_at = ? WHERE id = ? and org_id = ?`
	getRemoteWorkTypeQuery := `SELECT employee_id, type FROM remote_work_requests WHERE id = ? and org_id = ?`
	updateRemoteWorkBalanceQuery := `UPDATE remote_work_balances SET annual_balance = annual_balance - 1 WHERE org_id = ? and employee_id = ? and type = ?`

	_, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection", err)
			return
		}
	}(r.db)

	_, err = r.db.Exec(updateRemoteWorkBalanceQuery, status, updatedBy, time.Now(), requestID, orgID)
	if status == common.Approved.String() {
		var remoteWorkType string
		var employeeID string
		err = r.db.QueryRow(getRemoteWorkTypeQuery, requestID, orgID).Scan(&employeeID, &remoteWorkType)
		if err != nil {
			return err
		}
		_, err = r.db.Exec(updateRemoteWorkBalanceQuery, orgID, employeeID, remoteWorkType)
		if err != nil {
			return err
		}
	}
	_, err = r.db.Exec(statusUpdateQuery, status, updatedBy, time.Now(), requestID, orgID)
	return err
}

func (r *remoteWorkRepository) GetRemoteWorkRequestsInRange(orgID int, startDate, endDate time.Time) ([]model.RemoteWorkRequest, error) {
	query := `SELECT id, org_id, employee_id, type, start_date, end_date, reason, status, approved_by, created_at, updated_at FROM remote_work_requests WHERE org_id = ? AND start_date >= ? AND end_date <= ?`
	rows, err := r.db.Query(query, orgID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var requests []model.RemoteWorkRequest
	for rows.Next() {
		var request model.RemoteWorkRequest
		err := rows.Scan(&request.ID, &request.OrgID, &request.EmployeeID, &request.Type, &request.StartDate, &request.EndDate, &request.Reason, &request.Status, &request.UpdatedBy, &request.CreatedAt, &request.UpdatedAt)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}
