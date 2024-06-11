package util

import (
	"time"
)

// GetCurrentTimeInTimezone returns the current time in the specified timezone.
func GetCurrentTimeInTimezone(timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(location), nil
}

// ConvertToTimezone converts the given time to the specified timezone.
func ConvertToTimezone(t time.Time, timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(location), nil
}

// GetLocalTimezone fetches the local timezone dynamically.
func GetLocalTimezone() (string, error) {
	localLocation := time.Now().Location()
	return localLocation.String(), nil
}

// ExtractDate takes a time.Time object and returns a new time.Time object with only the date part
func ExtractDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
