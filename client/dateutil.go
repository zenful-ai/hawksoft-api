package hawksoft

import (
	"time"
)

// Common date/time formats used by the HawkSoft API
const (
	// DateTimeFormat is the format for date-time fields without timezone: "2024-10-08T00:00:00"
	DateTimeFormat = "2006-01-02T15:04:05"

	// DateFormat is the format for date-only fields: "2024-10-08"
	DateFormat = "2006-01-02"
)

// ParseDateTime parses a date-time string in the format "2024-10-08T00:00:00"
// Returns zero time and error if parsing fails.
func ParseDateTime(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}
	return time.Parse(DateTimeFormat, s)
}

// ParseDate parses a date string in the format "2024-10-08"
// Returns zero time and error if parsing fails.
func ParseDate(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}
	return time.Parse(DateFormat, s)
}

// FormatDateTime formats a time.Time to the API's date-time format: "2024-10-08T00:00:00"
func FormatDateTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(DateTimeFormat)
}

// FormatDate formats a time.Time to the API's date format: "2024-10-08"
func FormatDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(DateFormat)
}

// Helper functions for pointer conversions

// StringPtr returns a pointer to the given string
func StringPtr(s string) *string {
	return &s
}

// TimeToDateTimePtr converts a time.Time to a formatted date-time string pointer
func TimeToDateTimePtr(t time.Time) *string {
	if t.IsZero() {
		return nil
	}
	s := FormatDateTime(t)
	return &s
}

// TimeToDatePtr converts a time.Time to a formatted date string pointer
func TimeToDatePtr(t time.Time) *string {
	if t.IsZero() {
		return nil
	}
	s := FormatDate(t)
	return &s
}

// DateTimePtrToTime parses a date-time string pointer to time.Time
// Returns zero time if pointer is nil or parsing fails.
func DateTimePtrToTime(s *string) time.Time {
	if s == nil || *s == "" {
		return time.Time{}
	}
	t, _ := ParseDateTime(*s)
	return t
}

// DatePtrToTime parses a date string pointer to time.Time
// Returns zero time if pointer is nil or parsing fails.
func DatePtrToTime(s *string) time.Time {
	if s == nil || *s == "" {
		return time.Time{}
	}
	t, _ := ParseDate(*s)
	return t
}
