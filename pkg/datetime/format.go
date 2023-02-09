package datetime

import "time"

func ToISO8601String(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
