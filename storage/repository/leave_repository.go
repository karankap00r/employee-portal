package repository

import (
	"database/sql"
	"github.com/karankap00r/employee_portal/common"
	"github.com/karankap00r/employee_portal/storage/model"
	"time"
)

type LeaveRepository interface {
	GetLeaveBalance(orgID int, employeeID string) (*model.LeaveBalance, error)
	CreateLeaveRequest(request model.LeaveRequest) error
	UpdateLeaveRequestStatus(orgID, leaveRequestID int, status, approvedBy string) error
	GetLeavesInRange(orgID int, startDate, endDate time.Time) ([]model.LeaveRequest, error)
}

type leaveRepository struct {
	db *sql.DB
}

func NewLeaveRepository(db *sql.DB) LeaveRepository {
	return &leaveRepository{db}
}

func (r *leaveRepository) GetLeaveBalance(orgID int, employeeID string) (*model.LeaveBalance, error) {
	query := `SELECT id, org_id, employee_id, leave_type, annual_balance, created_at, updated_at FROM leave_balances WHERE org_id = ? AND employee_id = ?`
	row := r.db.QueryRow(query, orgID, employeeID)
	var balance model.LeaveBalance
	err := row.Scan(&balance.ID, &balance.OrgID, &balance.EmployeeID, &balance.LeaveType, &balance.AnnualBalance, &balance.CreatedAt, &balance.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (r *leaveRepository) CreateLeaveRequest(request model.LeaveRequest) error {
	query := `INSERT INTO leave_requests (org_id, employee_id, leave_category, start_date, end_date, reason, status, updated_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, request.OrgID, request.EmployeeID, request.LeaveCategory, request.StartDate, request.EndDate, request.Reason, request.Status, request.UpdatedBy, request.CreatedAt, request.UpdatedAt)
	return err
}

func (r *leaveRepository) UpdateLeaveRequestStatus(orgID, leaveRequestID int, status, updatedBy string) error {
	updateStatusQuery := `UPDATE leave_requests SET status = ?, updated_by = ?, updated_at = ? WHERE id = ? and org_id = ?`
	getLeaveTypeQuery := `SELECT employee_id, leave_category FROM leave_requests WHERE id = ? and org_id = ?`
	updateLeaveBalanceQuery := `UPDATE leave_balances SET annual_balance = annual_balance - 1 WHERE org_id = ? and employee_id = ? and leave_type = ?`

	tx, err := r.db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		return err
	}

	_, err = r.db.Exec(updateStatusQuery, status, updatedBy, time.Now(), leaveRequestID, orgID)
	if status == common.Approved.String() {
		var leaveType string
		var employeeID string
		err = r.db.QueryRow(getLeaveTypeQuery, leaveRequestID, orgID).Scan(&employeeID, &leaveType)
		if err != nil {
			return err
		}
		_, err = r.db.Exec(updateLeaveBalanceQuery, orgID, employeeID, leaveType)
		if err != nil {
			return err
		}
	}
	return err
}

func (r *leaveRepository) GetLeavesInRange(orgID int, startDate, endDate time.Time) ([]model.LeaveRequest, error) {
	query := `SELECT id, org_id, employee_id, leave_category, start_date, end_date, reason, status, updated_by, created_at, updated_at FROM leave_requests WHERE org_id = ? AND start_date >= ? AND end_date <= ?`
	rows, err := r.db.Query(query, orgID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var requests []model.LeaveRequest
	for rows.Next() {
		var request model.LeaveRequest
		err := rows.Scan(&request.ID, &request.OrgID, &request.EmployeeID, &request.LeaveCategory, &request.StartDate, &request.EndDate, &request.Reason, &request.Status, &request.UpdatedBy, &request.CreatedAt, &request.UpdatedAt)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}
