package datetime

import (
	"testing"
)

func Test_ParseUnixTime(t *testing.T) {
	unixTime := int64(1612872840)
	tm := ParseUnixTime(unixTime)
	if tm.IsZero() {
		t.Errorf("ParseJSDate(%d) = %v", unixTime, tm)
	}
}

func Test_ParseJSDate(t *testing.T) {
	jsDate := "Thu Feb 09 2023 15:54:00 GMT+0700 (Indochina Time)"
	tm, err := ParseJSDate(jsDate)
	if err != nil {
		t.Errorf("ParseJSDate(%s) error: %v", jsDate, err)
	}
	if tm.IsZero() {
		t.Errorf("ParseJSDate(%s) = %v", jsDate, tm)
	}
}
