package datetime

import "time"

func ParseUnixTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}
