package rtime

import "time"

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday time.Weekday

// Returns specifies a day of the week (Sunday = 0, ...).
func (t Time) Weekday() Weekday {
	return Weekday(time.Time(t).Weekday())
}

// Returns string representation ofWeekday
func (w Weekday) String() (weekday string) {
	switch w {
	case 0:
		weekday = "Воскресенье"
	case 1:
		weekday = "Понедельник"
	case 2:
		weekday = "Вторник"
	case 3:
		weekday = "Среда"
	case 4:
		weekday = "Четверг"
	case 5:
		weekday = "Пятница"
	case 6:
		weekday = "Суббота"
	}
	return
}

// Returns string representation of Day of Month
func (t Time) DayString() (result string) {
	day := time.Time(t).Day()
	return dayString(day)
}

func dayString(day int) (result string) {
	switch day {
	case 1:
		result = "Первое"
	case 2:
		result = "Второе"
	case 3:
		result = "Третье"
	case 4:
		result = "Четвертое"
	case 5:
		result = "Пятое"
	case 6:
		result = "Шестое"
	case 7:
		result = "Седьмое"
	case 8:
		result = "Восьмое"
	case 9:
		result = "Девятое"
	case 10:
		result = "Десятое"
	case 11:
		result = "Одиннадцатое"
	case 12:
		result = "Двенадцатое"
	case 13:
		result = "Тринадцатое"
	case 14:
		result = "Четырнадцатое"
	case 15:
		result = "Пятнадцатое"
	case 16:
		result = "Шестнадцатое"
	case 17:
		result = "Семнадцатое"
	case 18:
		result = "Восемнадцатое"
	case 19:
		result = "Девятнадцатое"
	case 20:
		result = "Двадцатое"
	case 21:
		result = "Двадцать первое"
	case 22:
		result = "Двадцать второе"
	case 23:
		result = "Двадцать третье"
	case 24:
		result = "Двадцать четвертое"
	case 25:
		result = "Двадцать пятое"
	case 26:
		result = "Двадцать шестое"
	case 27:
		result = "Двадцать седьмое"
	case 28:
		result = "Двадцать восьмое"
	case 29:
		result = "Двадцать девятое"
	case 30:
		result = "Тридцатое"
	case 31:
		result = "Тридцать первое"
	}
	return
}
