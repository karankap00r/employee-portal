package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/karankap00r/employee_portal/common"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/storage/repository"
)

const publicHolidayAPIURL = "https://holidays.abstractapi.com/v1/"

// PublicHolidayService is an interface for the public holiday service
type PublicHolidayService interface {
	// SyncAllCountries syncs public holidays for all countries
	SyncAllCountries() error

	// SyncCountries syncs public holidays for the given countries
	SyncCountries(countries []string) error

	// GetAllPublicHolidays returns all public holidays
	GetAllPublicHolidays() ([]model.PublicHoliday, error)

	// AddPublicHoliday adds a public holiday
	AddPublicHoliday(holiday model.PublicHoliday) error

	// UpdatePublicHolidayStatus updates the status of a public holiday
	UpdatePublicHolidayStatus(id int, status string) error

	// GetPublicHolidaysForNext7Days returns the public holidays for the next 7 days
	GetPublicHolidaysForNext7Days(country string) ([]model.PublicHoliday, error)

	// SendPublicHolidayAlert sends an email alert for upcoming public holidays
	SendPublicHolidayAlert(email, country string) error
}

// publicHolidayService is a struct for the public holiday service
type publicHolidayService struct {
	repo   repository.PublicHolidayRepository
	apiKey string
}

// NewPublicHolidayService returns a new public holiday service
func NewPublicHolidayService(repo repository.PublicHolidayRepository, apiKey string) PublicHolidayService {
	return &publicHolidayService{repo, apiKey}
}

// SyncAllCountries syncs public holidays for all countries
func (s *publicHolidayService) SyncAllCountries() error {
	countries, err := s.fetchCountryList()
	if err != nil {
		return err
	}
	return s.SyncCountries(countries)
}

// SyncCountries syncs public holidays for the given countries
func (s *publicHolidayService) SyncCountries(countries []string) error {
	for _, country := range countries {
		holidays, err := s.fetchHolidaysByCountry(country)
		if err != nil {
			return err
		}
		for _, holiday := range holidays {
			err := s.repo.CreatePublicHoliday(holiday)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetAllPublicHolidays returns all public holidays
func (s *publicHolidayService) GetAllPublicHolidays() ([]model.PublicHoliday, error) {
	return s.repo.GetAllPublicHolidays()
}

// AddPublicHoliday adds a public holiday
func (s *publicHolidayService) AddPublicHoliday(holiday model.PublicHoliday) error {
	holiday.CreatedAt = time.Now()
	holiday.UpdatedAt = time.Now()
	return s.repo.CreatePublicHoliday(holiday)
}

// UpdatePublicHolidayStatus updates the status of a public holiday
func (s *publicHolidayService) UpdatePublicHolidayStatus(id int, status string) error {
	if status != common.Active.String() && status != common.Inactive.String() {
		return errors.New("invalid status")
	}
	return s.repo.UpdatePublicHolidayStatus(id, status)
}

// fetchCountryList fetches the list of countries
func (s *publicHolidayService) fetchCountryList() ([]string, error) {
	// Fetch country list from a public API or static list
	return []string{"US", "CA", "GB"}, nil // Example countries
}

// fetchHolidaysByCountry fetches the public holidays for a country
func (s *publicHolidayService) fetchHolidaysByCountry(country string) ([]model.PublicHoliday, error) {
	url := publicHolidayAPIURL + "?api_key=" + s.apiKey + "&country=" + country + "&year=" + time.Now().Format("2006")
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch public holidays")
	}

	var holidays []model.PublicHoliday
	err = json.NewDecoder(resp.Body).Decode(&holidays)
	if err != nil {
		return nil, err
	}
	return holidays, nil
}

// GetPublicHolidaysForNext7Days returns the public holidays for the next 7 days
func (s *publicHolidayService) GetPublicHolidaysForNext7Days(country string) ([]model.PublicHoliday, error) {
	return s.repo.GetPublicHolidaysForNext7Days(country)
}

// SendPublicHolidayAlert sends an email alert for upcoming public holidays
func (s *publicHolidayService) SendPublicHolidayAlert(email, country string) error {
	holidays, err := s.GetPublicHolidaysForNext7Days(country)
	if err != nil {
		return err
	}

	if len(holidays) == 0 {
		return nil
	}

	subject := "Upcoming Public Holidays"
	body := "Here are the upcoming public holidays in the next 7 days:\n\n"
	for _, holiday := range holidays {
		body += fmt.Sprint(holiday.Name, " on ", holiday.StartDate, "\n")
	}

	return sendEmail(email, subject, body)
}
