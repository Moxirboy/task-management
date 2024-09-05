package utils

import (
	"strings"
	"time"
)

const (
	dateTimeFormat = time.RFC3339
	dateFormat     = `2006-01-02`
)

func ParseDateTime(val string) *time.Time {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil
	}
	dateTime, _ := time.Parse(dateTimeFormat, val)

	return &dateTime
}

func ParseDate(val string) *time.Time {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil
	}
	date, _ := time.Parse(dateFormat, val)

	return &date
}

func DateFormat(date *time.Time) string {
	if date == nil {
		return ""
	}
	return date.Format(dateTimeFormat)
}
