package rtime

import "time"

func (t Time) Year() int {
	return time.Time(t).Year()
}

// Returns string representation of Year
func (t Time) YearString() (year string) {
	y := t.Year()
	switch {
	case y == 1900:
		year = "Одна тысяча девятисотый"
	case y < 2000 && y >= 1900:
		year = "Одна тысяча девятьсот " + lastDecadeYearString(y-1900)
	case y == 2000:
		year = "Двухтысячный"
	case y > 2000 && y < 3000:
		year = "Две тысячи " + lastDecadeYearString(y-2000)
	}
	return
}

// Returns string representation of Year in case
// Двухтысячного года, Одна тысяча девятьсот четертого года
func (t Time) YearStringInCase() (year string) {
	y := t.Year()
	switch {
	case y == 1900:
		year = "Одна тысяча девятисотого"
	case y < 2000 && y >= 1900:
		year = "Одна тысяча девятьсот " + lastDecadeYearStringInCase(y-1900)
	case y == 2000:
		year = "Двухтысячного"
	case y > 2000 && y < 3000:
		year = "Две тысячи " + lastDecadeYearStringInCase(y-2000)
	}
	return year + " года"
}

func lastDecadeYearString(y int) (result string) {
	switch {
	case y == 10:
		result = "десятого"
	case y == 11:
		result = "одиннадцатого"
	case y == 12:
		result = "двенадцатого"
	case y == 13:
		result = "тринадцатого"
	case y == 14:
		result = "четырнадцатого"
	case y == 15:
		result = "пятнадцатого"
	case y == 16:
		result = "шестнадцатого"
	case y == 17:
		result = "семьнадцатого"
	case y == 18:
		result = "восемьнадцатого"
	case y == 19:
		result = "девятьнадцатого"
	case y == 20:
		result = "двадцатого"
	case y == 30:
		result = "тридцатого"
	case y == 40:
		result = "сорокового"
	case y == 50:
		result = "пятидесятого"
	case y == 60:
		result = "шестидесятого"
	case y == 70:
		result = "семидесятого"
	case y == 80:
		result = "восьмидесятого"
	case y == 90:
		result = "девяностого"
	default:
		d := y % 10
		decade := y - d
		result = decadeString(decade)
		result += " " + lastYearDayString(d)
	}
	return
}

func lastDecadeYearStringInCase(y int) (result string) {
	switch {
	case y == 10:
		result = "десятого"
	case y == 11:
		result = "одиннадцатого"
	case y == 12:
		result = "двенадцатого"
	case y == 13:
		result = "тринадцатого"
	case y == 14:
		result = "четырнадцатого"
	case y == 15:
		result = "пятнадцатого"
	case y == 16:
		result = "шестнадцатого"
	case y == 17:
		result = "семьнадцатого"
	case y == 18:
		result = "восемьнадцатого"
	case y == 19:
		result = "девятьнадцатого"
	case y == 20:
		result = "двадцатого"
	case y == 30:
		result = "тридцатого"
	case y == 40:
		result = "сороковой"
	case y == 50:
		result = "пятидесятого"
	case y == 60:
		result = "шестидесятого"
	case y == 70:
		result = "семидесятого"
	case y == 80:
		result = "восьмидесятого"
	case y == 90:
		result = "девяностого"
	default:
		d := y % 10
		decade := y - d
		result = decadeString(decade)
		result += " " + lastYearDayStringInCase(d)
	}
	return
}

func decadeString(y int) (result string) {
	switch y {
	case 20:
		result = "двадцать"
	case 30:
		result = "тридцать"
	case 40:
		result = "сорок"
	case 50:
		result = "пятьдесят"
	case 60:
		result = "шестьдесят"
	case 70:
		result = "семьдесят"
	case 80:
		result = "восемьдесят"
	case 90:
		result = "девяносто"
	}
	return
}

func lastYearDayString(y int) (result string) {
	switch y {
	case 1:
		result = "первый"
	case 2:
		result = "второй"
	case 3:
		result = "третий"
	case 4:
		result = "четвертый"
	case 5:
		result = "пятый"
	case 6:
		result = "шестой"
	case 7:
		result = "седьмой"
	case 8:
		result = "восьмой"
	case 9:
		result = "девятый"
	}
	return
}

func lastYearDayStringInCase(y int) (result string) {
	switch y {
	case 1:
		result = "первого"
	case 2:
		result = "второго"
	case 3:
		result = "третьего"
	case 4:
		result = "четвертого"
	case 5:
		result = "пятого"
	case 6:
		result = "шестого"
	case 7:
		result = "седьмого"
	case 8:
		result = "восьмого"
	case 9:
		result = "девятого"
	}
	return
}
