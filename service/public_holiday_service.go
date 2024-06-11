package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/karankap00r/employee_portal/common"
	"github.com/karankap00r/employee_portal/storage/model"
	"github.com/karankap00r/employee_portal/storage/repository"
)

const publicHolidayAPIURL = "https://holidays.abstractapi.com/v1/"

type PublicHolidayService interface {
	SyncAllCountries() error
	SyncCountries(countries []string) error
	GetAllPublicHolidays() ([]model.PublicHoliday, error)
	AddPublicHoliday(holiday model.PublicHoliday) error
	UpdatePublicHolidayStatus(id int, status string) error
}

type publicHolidayService struct {
	repo   repository.PublicHolidayRepository
	apiKey string
}

func NewPublicHolidayService(repo repository.PublicHolidayRepository, apiKey string) PublicHolidayService {
	return &publicHolidayService{repo, apiKey}
}

func (s *publicHolidayService) SyncAllCountries() error {
	countries, err := s.fetchCountryList()
	if err != nil {
		return err
	}
	return s.SyncCountries(countries)
}

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

func (s *publicHolidayService) GetAllPublicHolidays() ([]model.PublicHoliday, error) {
	return s.repo.GetAllPublicHolidays()
}

func (s *publicHolidayService) AddPublicHoliday(holiday model.PublicHoliday) error {
	holiday.CreatedAt = time.Now()
	holiday.UpdatedAt = time.Now()
	return s.repo.CreatePublicHoliday(holiday)
}

func (s *publicHolidayService) UpdatePublicHolidayStatus(id int, status string) error {
	if status != common.Active.String() && status != common.Inactive.String() {
		return errors.New("invalid status")
	}
	return s.repo.UpdatePublicHolidayStatus(id, status)
}

func (s *publicHolidayService) fetchCountryList() ([]string, error) {
	// Fetch country list from a public API or static list
	return []string{"US", "CA", "GB"}, nil // Example countries
}

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
