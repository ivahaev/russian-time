package rtime

import (
	"strings"
	"time"
)

// A Duration represents the elapsed time between two instants as an int64 nanosecond count.
// The representation limits the largest representable duration to approximately 290 years.
type Duration time.Duration

// Hours returns the duration as a floating point number of hours.
func (d Duration) Hours() float64 {
	return time.Duration(d).Hours()
}

// Minutes returns the duration as a floating point number of minutes.
func (d Duration) Minutes() float64 {
	return time.Duration(d).Minutes()
}

// Nanoseconds returns the duration as an integer nanosecond count.
func (d Duration) Nanoseconds() int64 {
	return time.Duration(d).Nanoseconds()
}

// Seconds returns the duration as a floating point number of seconds.
func (d Duration) Seconds() float64 {
	return time.Duration(d).Seconds()
}

// Returns a string representation of the duration in russian language
func (d Duration) String() (result string) {
	var seconds, minutes, hours int
	seconds = int(d.Seconds())
	if seconds > 60 {
		minutes = (seconds - seconds%60) / 60
		seconds = seconds % 60
	}
	if minutes > 59 {
		hours = (minutes - minutes%60) / 60
		minutes = minutes - hours*60
		result = numberInString(hours, false)
		result += " " + hoursTail(hours)
	}
	if minutes != 0 {
		if result != "" {
			result += ", "
		}
		result += strings.ToLower(numberInString(minutes, true))
		result += " " + minutesTail(minutes)
	}
	if seconds != 0 {
		if result != "" {
			result += ", "
		}
		result += strings.ToLower(numberInString(seconds, true))
		result += " " + secondsTail(seconds)

	}
	return
}

// Returns a string representation of the approximate duration in russian language
func (d Duration) StringApproximate() (result string) {
	var seconds, minutes, hours, days, months, years int
	seconds = int(d.Seconds())
	if seconds > 60 {
		minutes = int(d.Minutes())
	}
	if minutes > 59 {
		hours = int(d.Hours())
		minutes = minutes - hours*60
	}
	if hours > 24 {
		days = (hours - hours%24) / 24
		hours = hours - days*24
	}
	if days > 365 {
		years = (days - days%365) / 365
		days = days - years*365
	}
	if days > 30 {
		months = (days - days%30) / 30
		days = days - months*30
	}
	if years > 0 {
		if months < 3 {
			result = numberInString(years, false) + " " + yearsTail(years)
		} else {
			result = "Более"
			if years > 1 {
				result = " " + strings.ToLower(numberStringInGenitiveCase(years, false))
			}
			result += " " + strings.ToLower(numberStringInGenitiveCase(years, false)) + " " + strings.ToLower(yearsTailInGenitiveCase(years))
		}
	} else if months > 0 {
		if days < 8 {
			result = numberInString(months, false) + " " + monthsTail(months)
		} else {
			result = "Более"
			if months > 1 {
				result = " " + strings.ToLower(numberStringInGenitiveCase(months, false))
			}
			result += " " + strings.ToLower(numberStringInGenitiveCase(months, false)) + " " + strings.ToLower(monthsTailInGenitiveCase(months))
		}
	} else if days > 0 {
		if hours < 5 {
			result = numberInString(days, false) + " " + daysTail(days)
		} else {
			result = "Более "
			if days == 1 {
				result += "суток"
			} else {
				result += strings.ToLower(numberStringInGenitiveCase(days, false)) + " суток"
			}
		}
	} else if hours > 0 {
		if minutes < 16 {
			result = numberInString(hours, false) + " " + hoursTail(hours)
		} else {
			result = "Более "
			if hours == 1 {
				result += "часа"
			} else {
				result += strings.ToLower(numberStringInGenitiveCase(hours, false))
				result += " " + strings.ToLower(hoursTailInGenitiveCase(hours))
			}
		}
	} else if minutes > 0 {
		if minutes == 1 {
			result = "Минуту"
		} else {
			result = numberInString(minutes, true) + " " + minutesTail(minutes)
		}
	} else {
		result = "Менее минуты"
	}
	result += " назад"
	return
}

func hoursTailInGenitiveCase(hours int) (result string) {
	switch {
	case hours > 20 && hours < 100:
		hours = hours % 10
	case hours > 100:
		hours = hours % 100 % 10
	}
	switch hours {
	case 1:
		result = "часа"
	default:
		result = "часов"
	}
	return
}

func yearsTail(years int) (result string) {
	switch {
	case years > 20 && years < 100:
		years = years % 10
	case years > 100:
		years = years % 100 % 10
	}
	switch years {
	case 1:
		result = "год"
	case 2, 3, 4:
		result = "года"
	default:
		result = "лет"
	}
	return
}

func yearsTailInGenitiveCase(years int) (result string) {
	switch {
	case years > 20 && years < 100:
		years = years % 10
	case years > 100:
		years = years % 100 % 10
	}
	switch years {
	case 1:
		result = "года"
	default:
		result = "лет"
	}
	return
}

func monthsTail(months int) (result string) {
	switch {
	case months > 20 && months < 100:
		months = months % 10
	case months > 100:
		months = months % 100 % 10
	}
	switch months {
	case 1:
		result = "месяц"
	case 2, 3, 4:
		result = "месяца"
	default:
		result = "месяцев"
	}
	return
}

func monthsTailInGenitiveCase(months int) (result string) {
	switch {
	case months > 20 && months < 100:
		months = months % 10
	case months > 100:
		months = months % 100 % 10
	}
	switch months {
	case 1:
		result = "месяца"
	default:
		result = "месяцев"
	}
	return
}

func daysTail(days int) (result string) {
	switch {
	case days > 20 && days < 100:
		days = days % 10
	case days > 100:
		days = days % 100 % 10
	}
	switch days {
	case 1:
		result = "день"
	case 2, 3, 4:
		result = "дня"
	default:
		result = "дней"
	}
	return
}

func daysTailInGenitiveCase(days int) (result string) {
	switch {
	case days > 20 && days < 100:
		days = days % 10
	case days > 100:
		days = days % 100 % 10
	}
	switch days {
	case 1:
		result = "дня"
	default:
		result = "дней"
	}
	return
}

func minutesTailInGenitiveCase(minutes int) (result string) {
	switch {
	case minutes > 20 && minutes < 60:
		minutes = minutes % 10
	case minutes > 60:
		minutes = minutes % 60 % 10
	}
	switch minutes {
	case 1:
		result = "минуты"
	default:
		result = "минут"
	}
	return
}

func secondsTailInGenitiveCase(seconds int) (result string) {
	switch {
	case seconds > 20 && seconds < 60:
		seconds = seconds % 10
	case seconds > 60:
		seconds = seconds % 60 % 10
	}
	switch seconds {
	case 1:
		result = "секунды"
	default:
		result = "секунд"
	}
	return
}

func numberStringInGenitiveCase(num int, female bool) (result string) {
	switch num {
	case 0:
		result = "Нуля"
	case 1:
		if female {
			result = "Одной"
		} else {
			result = "Одного"
		}
	case 2:
		result = "Двух"
	case 3:
		result = "Трёх"
	case 4:
		result = "Четырёх"
	case 5:
		result = "Пяти"
	case 6:
		result = "Шести"
	case 7:
		result = "Семи"
	case 8:
		result = "Восьми"
	case 9:
		result = "Девяти"
	case 10:
		result = "Десяти"
	case 11:
		result = "Одиннадцати"
	case 12:
		result = "Двенадцати"
	case 13:
		result = "Тринадцати"
	case 14:
		result = "Четырнадцати"
	case 15:
		result = "Пятнадцати"
	case 16:
		result = "Шестнадцати"
	case 17:
		result = "Семнадцати"
	case 18:
		result = "Восемнадцати"
	case 19:
		result = "Девятнадцати"
	case 20:
		result = "Двадцати"
	case 30:
		result = "Тридцати"
	case 40:
		result = "Сорока"
	case 50:
		result = "Пятидесяти"
	case 60:
		result = "Шестидесяти"
	case 70:
		result = "Семидесяти"
	case 80:
		result = "Восьмидесяти"
	case 90:
		result = "Девяноста"
	default:
		lastDigit := num % 10
		result = numberInString(num-lastDigit, female)
		result += " " + strings.ToLower(numberInString(lastDigit, female))
	}
	return
}
