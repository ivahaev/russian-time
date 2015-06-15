# russian-time

A tiny library to represent DateTime in russian language

Installation:
--
    go get "github.com/ivahaev/russian-time"


## Usage

```go
import (
    "fmt"
    "github.com/ivahaev/russian-time"
    "time"
)
// Get current local DateTime
t := rtime.Now()
// Or if you are using time.Time object:
standardTime := time.Now()
t = rtime.Time(standardTime)

// Get full string representation:
fmt.Println(t.String())
// Шестнадцатое июня две тысячи пятнадцатого года, один час, тридцать шесть минут, тридцать три секунды

// Get date string representation:
fmt.Println(t.DateString())
// Шестнадцатое июня две тысячи пятнадцатого года

// Get time string representation:
fmt.Println(t.TimeString())
// Один час, тридцать шесть минут

// Get time with seconds string representation:
fmt.Println(t.TimeStringWithSeconds())
// Один час, тридцать шесть минут, тридцать три секунды
```


#### type Month

```go
type Month time.Month
```

Base type for month

#### func (Month) String

```go
func (m Month) String() (month string)
```
Returns string representation of Month

#### func (Month) StringInCase

```go
func (m Month) StringInCase() (month string)
```
Returns representation of Month in case:
Января, Декабря...

#### type Time

```go
type Time time.Time
```

Base type for DateTime. Can be converted from time.Time

#### func  Now

```go
func Now() Time
```
Returns current DateTime

#### func (Time) DateString

```go
func (t Time) DateString() (result string)
```
Returns date string representation of Time

#### func (Time) DayString

```go
func (t Time) DayString() (result string)
```
Returns string representation of Day of Month

#### func (Time) Hour

```go
func (t Time) Hour() int
```
Returns the hour within the day specified by t, in the range [0, 23].

#### func (Time) HourString

```go
func (t Time) HourString() (result string)
```
Returns string representation of Hour

#### func (Time) Minute

```go
func (t Time) Minute() int
```
Returns the minute offset within the hour specified by t, in the range [0, 59].

#### func (Time) MinuteString

```go
func (t Time) MinuteString() (result string)
```
Returns string representation of Minute

#### func (Time) Month

```go
func (t Time) Month() Month
```
Returns int month [0,12]

#### func (Time) Second

```go
func (t Time) Second() int
```
Returns the second offset within the minute specified by t, 
in the range [0,59].

#### func (Time) SecondString

```go
func (t Time) SecondString() (result string)
```
Returns string representation of Second

#### func (Time) String

```go
func (t Time) String() (result string)
```
Returns full string representation of Time

#### func (Time) TimeString

```go
func (t Time) TimeString() (result string)
```
Returns time string representation of Time without seconds

#### func (Time) TimeStringWithSeconds

```go
func (t Time) TimeStringWithSeconds() (result string)
```
Returns time string representation of Time with seconds

#### func (Time) Weekday

```go
func (t Time) Weekday() Weekday
```
Returns specifies a day of the week (Sunday = 0, ...).

#### func (Time) Year

```go
func (t Time) Year() int
```

#### func (Time) YearString

```go
func (t Time) YearString() (year string)
```
Returns string representation of Year

#### func (Time) YearStringInCase

```go
func (t Time) YearStringInCase() (year string)
```
Returns string representation of Year in case:
Двухтысячного года, 
Одна тысяча девятьсот четертого года

#### type Weekday

```go
type Weekday time.Weekday
```

A Weekday specifies a day of the week (Sunday = 0, ...).

#### func (Weekday) String

```go
func (w Weekday) String() (weekday string)
```
Returns string representation ofWeekday
