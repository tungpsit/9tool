package datetime

import (
	"regexp"
	"time"
)

const (
	jsDateLayout = "Mon Jan 02 2006 15:04:05 GMT-0700"
)

// Parse unix time to time.Time
// Example: 1612872840
func ParseUnixTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}

// Parse javascript date string to time.Time
// Example: "Thu Feb 09 2023 15:54:00 GMT+0700 (Indochina Time)"
func ParseJSDate(jsDate string) (time.Time, error) {
	re := regexp.MustCompile(`\s\(.*?\)$`)
	input := re.ReplaceAllString(jsDate, "")
	return time.Parse(jsDateLayout, input)
}
