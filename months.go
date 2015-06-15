package rtime

import "time"

// Base type for month
type Month time.Month

// Returns int month [0,12]
func (t Time) Month() Month {
	return Month(time.Time(t).Month())
}

// Returns string representation of Month
func (m Month) String() (month string) {
	switch m {
	case Month(time.January):
		month = "Январь"
	case Month(time.February):
		month = "Февраль"
	case Month(time.March):
		month = "Март"
	case Month(time.April):
		month = "Апрель"
	case Month(time.May):
		month = "Май"
	case Month(time.June):
		month = "Июнь"
	case Month(time.July):
		month = "Июль"
	case Month(time.August):
		month = "Август"
	case Month(time.September):
		month = "Сентябрь"
	case Month(time.October):
		month = "Октябрь"
	case Month(time.November):
		month = "Ноябрь"
	case Month(time.December):
		month = "Декабрь"
	}
	return
}

// Returns representation of Month in case
// Января, Декабря...
func (m Month) StringInCase() (month string) {
	switch m {
	case Month(time.January):
		month = "Января"
	case Month(time.February):
		month = "Февраля"
	case Month(time.March):
		month = "Марта"
	case Month(time.April):
		month = "Апреля"
	case Month(time.May):
		month = "Мая"
	case Month(time.June):
		month = "Июня"
	case Month(time.July):
		month = "Июля"
	case Month(time.August):
		month = "Августа"
	case Month(time.September):
		month = "Сентября"
	case Month(time.October):
		month = "Октября"
	case Month(time.November):
		month = "Ноября"
	case Month(time.December):
		month = "Декабря"
	}
	return
}
