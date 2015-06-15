package rtime

import (
	"strings"
	"time"
)

type Time time.Time

func Now() Time {
	return Time(time.Now())
}

func (t Time) String() (result string) {
	result = t.DayString()
	result += " " + strings.ToLower(t.Month().StringInCase())
	result += " " + strings.ToLower(t.YearStringInCase())
	return
}
