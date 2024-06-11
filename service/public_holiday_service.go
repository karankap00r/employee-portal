package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
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
	GetPublicHolidaysForNext7Days(country string) ([]model.PublicHoliday, error)
	SendPublicHolidayAlert(email, country string) error
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

func (s *publicHolidayService) GetPublicHolidaysForNext7Days(country string) ([]model.PublicHoliday, error) {
	return s.repo.GetPublicHolidaysForNext7Days(country)
}

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

func sendEmail(to, subject, body string) error {
	from := "your-email@example.com"
	password := "your-email-password"

	smtpHost := "smtp.example.com"
	smtpPort := "587"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" + body

	auth := smtp.PlainAuth("", from, password, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}
