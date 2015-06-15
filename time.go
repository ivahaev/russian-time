package rtime

import (
	"strings"
	"time"
)

// Returns the minute offset within the hour specified by t, in the range [0, 59].
func (t Time) Minute() int {
	return time.Time(t).Minute()
}

// Returns the second offset within the minute specified by t, in the range [0, 59].
func (t Time) Second() int {
	return time.Time(t).Second()
}

// Returns the hour within the day specified by t, in the range [0, 23].
func (t Time) Hour() int {
	return time.Time(t).Hour()
}

// Returns string representation of Hour
func (t Time) HourString() (result string) {
	result = numberInString(t.Hour(), false)
	switch t.Hour() {
	case 1:
		result += " час"
	case 2, 3, 4:
		result += " часа"
	default:
		result += " часов"
	}
	return
}

// Returns string representation of Minute
func (t Time) MinuteString() (result string) {
	result = numberInString(t.Minute(), true)
	switch t.Minute() {
	case 1:
		result += " минута"
	case 2, 3, 4:
		result += " минуты"
	default:
		result += " минут"
	}
	return
}

// Returns string representation of Second
func (t Time) SecondString() (result string) {
	result = numberInString(t.Second(), true)
	switch t.Second() {
	case 1:
		result += " секунда"
	case 2, 3, 4:
		result += " секунды"
	default:
		result += " секунд"
	}
	return
}

func numberInString(num int, female bool) (result string) {
	switch num {
	case 0:
		result = "Ноль"
	case 1:
		if female {
			result = "Одна"
		} else {
			result = "Один"
		}
	case 2:
		if female {
			result = "Две"
		} else {
			result = "Две"
		}
	case 3:
		result = "Три"
	case 4:
		result = "Четыре"
	case 5:
		result = "Пять"
	case 6:
		result = "Шесть"
	case 7:
		result = "Семь"
	case 8:
		result = "Восемь"
	case 9:
		result = "Девять"
	case 10:
		result = "Десять"
	case 11:
		result = "Одиннадцать"
	case 12:
		result = "Двенадцать"
	case 13:
		result = "Тринадцать"
	case 14:
		result = "Четырнадцать"
	case 15:
		result = "Пятнадцать"
	case 16:
		result = "Шестнадцать"
	case 17:
		result = "Семнадцать"
	case 18:
		result = "Восемнадцать"
	case 19:
		result = "Девятнадцать"
	case 20:
		result = "Двадцать"
	case 30:
		result = "Тридцать"
	case 40:
		result = "Сорок"
	case 50:
		result = "Пятьдесят"
	case 60:
		result = "Шестьдесят"
	case 70:
		result = "Семьдесят"
	case 80:
		result = "Восемьдесят"
	case 90:
		result = "Девяносто"
	default:
		lastDigit := num % 10
		result = numberInString(num-lastDigit, female)
		result += " " + strings.ToLower(numberInString(lastDigit, female))
	}
	return
}