package rtime

import (
	"strings"
	"time"
)

// Base type for DateTime. Can be converted from time.Time
type Time time.Time

// Returns current DateTime
func Now() Time {
	return Time(time.Now())
}

// Returns full string representation of Time
func (t Time) String() (result string) {
	result = t.DateString()
	result += ", " + strings.ToLower(t.TimeStringWithSeconds())
	return
}

// Returns date string representation of Time
func (t Time) DateString() (result string) {
	result = t.DayString()
	result += " " + strings.ToLower(t.Month().StringInCase())
	result += " " + strings.ToLower(t.YearStringInCase())
	return
}

// Returns time string representation of Time without seconds
func (t Time) TimeString() (result string) {
	result = t.HourString()
	result += ", " + strings.ToLower(t.MinuteString())
	return
}

// Returns time string representation of Time with seconds
func (t Time) TimeStringWithSeconds() (result string) {
	result = t.HourString()
	result += ", " + strings.ToLower(t.MinuteString())
	result += ", " + strings.ToLower(t.SecondString())
	return
}
