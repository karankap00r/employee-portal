package repository

import (
	"database/sql"
	"time"

	"github.com/karankap00r/employee_portal/storage/model"
)

type PublicHolidayRepository interface {
	CreatePublicHoliday(holiday model.PublicHoliday) error
	UpdatePublicHolidayStatus(id int, status string) error
	GetAllPublicHolidays() ([]model.PublicHoliday, error)
	GetPublicHolidaysByCountry(country string) ([]model.PublicHoliday, error)
}

type publicHolidayRepository struct {
	db *sql.DB
}

func NewPublicHolidayRepository(db *sql.DB) PublicHolidayRepository {
	return &publicHolidayRepository{db}
}

func (r *publicHolidayRepository) CreatePublicHoliday(holiday model.PublicHoliday) error {
	query := `INSERT INTO public_holidays (country, start_date, end_date, name, is_mandatory, status, created_by, created_at, updated_at)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, holiday.Country, holiday.StartDate, holiday.EndDate, holiday.Name, holiday.IsMandatory, holiday.Status, holiday.CreatedBy, holiday.CreatedAt, holiday.UpdatedAt)
	return err
}

func (r *publicHolidayRepository) UpdatePublicHolidayStatus(id int, status string) error {
	query := `UPDATE public_holidays SET status = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.Exec(query, status, time.Now(), id)
	return err
}

func (r *publicHolidayRepository) GetAllPublicHolidays() ([]model.PublicHoliday, error) {
	query := `SELECT id, country, start_date, end_date, name, is_mandatory, status, created_by, created_at, updated_at FROM public_holidays`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var holidays []model.PublicHoliday
	for rows.Next() {
		var holiday model.PublicHoliday
		err := rows.Scan(&holiday.ID, &holiday.Country, &holiday.StartDate, &holiday.EndDate, &holiday.Name, &holiday.IsMandatory, &holiday.Status, &holiday.CreatedBy, &holiday.CreatedAt, &holiday.UpdatedAt)
		if err != nil {
			return nil, err
		}
		holidays = append(holidays, holiday)
	}
	return holidays, nil
}

func (r *publicHolidayRepository) GetPublicHolidaysByCountry(country string) ([]model.PublicHoliday, error) {
	query := `SELECT id, country, start_date, end_date, name, is_mandatory, status, created_by, created_at, updated_at FROM public_holidays WHERE country = ?`
	rows, err := r.db.Query(query, country)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var holidays []model.PublicHoliday
	for rows.Next() {
		var holiday model.PublicHoliday
		err := rows.Scan(&holiday.ID, &holiday.Country, &holiday.StartDate, &holiday.EndDate, &holiday.Name, &holiday.IsMandatory, &holiday.Status, &holiday.CreatedBy, &holiday.CreatedAt, &holiday.UpdatedAt)
		if err != nil {
			return nil, err
		}
		holidays = append(holidays, holiday)
	}
	return holidays, nil
}
