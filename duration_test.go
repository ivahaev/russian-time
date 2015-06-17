package rtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var testStringVals = []testPair{
	{
		[]string{"2012-11-01T22:08:41+00:00", "2012-11-01T23:08:41+00:00"},
		"Один час",
	},
	{
		[]string{"2012-11-01T21:08:41+00:00", "2012-11-01T23:10:41+00:00"},
		"Два часа, две минуты",
	},
	{
		[]string{"2012-11-01T20:08:41+00:00", "2012-11-01T23:09:41+00:00"},
		"Три часа, одна минута",
	},
	{
		[]string{"2012-11-01T19:08:41+00:00", "2012-11-01T23:11:41+00:00"},
		"Четыре часа, три минуты",
	},
	{
		[]string{"2012-11-01T18:08:41+00:00", "2012-11-01T23:12:42+00:00"},
		"Пять часов, четыре минуты, одна секунда",
	},
	{
		[]string{"2012-11-01T10:45:00+00:00", "2012-11-01T23:12:43+00:00"},
		"Двенадцать часов, двадцать семь минут, сорок три секунды",
	},
}

var testStringApproximateVals = []testPair{
	{
		[]string{"2012-11-01T22:08:41+00:00", "2012-11-01T23:08:41+00:00"},
		"Один час назад",
	},
	{
		[]string{"2012-11-01T21:08:41+00:00", "2012-11-01T23:10:41+00:00"},
		"Два часа назад",
	},
	{
		[]string{"2012-11-01T20:08:41+00:00", "2012-11-01T23:29:41+00:00"},
		"Более трёх часов назад",
	},
	{
		[]string{"2012-11-01T00:00:00+00:00", "2012-11-01T00:00:30+00:00"},
		"Менее минуты назад",
	},
	{
		[]string{"2012-11-01T00:00:00+00:00", "2012-11-01T00:01:20+00:00"},
		"Минуту назад",
	},
	{
		[]string{"2012-11-01T00:00:00+00:00", "2012-11-01T00:05:59+00:00"},
		"Пять минут назад",
	},
	{
		[]string{"2012-11-01T00:00:00+00:00", "2012-11-11T12:05:59+00:00"},
		"Более десяти суток назад",
	},
	{
		[]string{"2012-11-01T00:00:00+00:00", "2012-12-03T12:05:59+00:00"},
		"Один месяц назад",
	},
	{
		[]string{"2012-11-01T00:00:00+00:00", "2012-12-11T12:05:59+00:00"},
		"Более одного месяца назад",
	},
	{
		[]string{"2012-11-01T00:00:00+00:00", "2013-11-11T12:05:59+00:00"},
		"Один год назад",
	},
	{
		[]string{"2012-11-01T00:00:00+00:00", "2014-03-11T12:05:59+00:00"},
		"Более одного года назад",
	},
}

type testPair struct {
	times  []string
	result string
}

func init() {

}

func TestString(t *testing.T) {
	for _, val := range testStringVals {
		t1, _ := time.Parse(time.RFC3339, val.times[0])
		t2, _ := time.Parse(time.RFC3339, val.times[1])
		duration := Duration(t2.Sub(t1))
		str := duration.String()
		assert.Equal(t, val.result, str, "String should equal")
	}
}

func TestStringApproximate(t *testing.T) {
	for _, val := range testStringApproximateVals {
		t1, _ := time.Parse(time.RFC3339, val.times[0])
		t2, _ := time.Parse(time.RFC3339, val.times[1])
		duration := Duration(t2.Sub(t1))
		str := duration.StringApproximate()
		assert.Equal(t, val.result, str, "String should equal")
	}
}
