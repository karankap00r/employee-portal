package cron

import (
	"github.com/karankap00r/employee_portal/service"
	"log"
	"net/smtp"
	"time"

	_ "github.com/robfig/cron/v3"
)

func StartCronJob(emailReportService service.EmailReportService, publicHolidayService service.PublicHolidayService) {
	c := cron.New()

	// Fetch active email reports for WeeklyHolidaysReport
	reports, err := emailReportService.GetActiveEmailReports("WeeklyHolidaysReport")
	if err != nil {
		log.Fatalf("Failed to fetch active email reports: %v", err)
	}

	for _, report := range reports {
		report := report // capture range variable

		// Schedule based on cron_frequency
		_, err := c.AddFunc(report.CronFrequency, func() {
			// Convert to IST
			location, _ := time.LoadLocation("Asia/Kolkata")
			currentTime := time.Now().In(location)

			// If today is not Sunday, return
			if currentTime.Weekday() != time.Sunday {
				return
			}

			holidays, err := publicHolidayService.GetPublicHolidaysForNext7Days(report.Country)
			if err != nil {
				log.Printf("Failed to fetch public holidays for org %v: %v", report.OrgID, err)
				return
			}

			if len(holidays) == 0 {
				return
			}

			subject := "Upcoming Public Holidays"
			body := "Here are the upcoming public holidays in the next 7 days:\n\n"
			for _, holiday := range holidays {
				body += holiday.Name + " on " + holiday.StartDate + "\n"
			}

			err = sendEmail(report.Email, subject, body)
			if err != nil {
				log.Printf("Failed to send email to org %v: %v", report.OrgID, err)
			}
		})

		if err != nil {
			log.Fatalf("Failed to schedule cron job for org %v: %v", report.OrgID, err)
		}
	}

	c.Start()
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
