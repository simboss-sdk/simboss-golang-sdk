package time

import (
	"time"
	"fmt"
	)

const (
	timeFormat = "2006-01-02 15:04:05"
	dateFormat = "2006-01-02"
)

func Now() Time {
	t := time.Now()
	return Time(t)
}

func Today() Date {
	t := time.Now()
	return Date(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0,0, time.Local))
}

func NewTime(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return Time(time.Date(year, month, day, hour, min, sec, nsec, loc))
}

func NewDate(year int, month time.Month, day int, loc *time.Location) Date {
	return Date(time.Date(year, month, day, 0, 0, 0,0, time.Local))
}

func ParseTime(timeString string) (Time, error) {
	t, err := time.Parse(timeFormat, timeString)
	if err != nil {
		return Time(time.Time{}), err
	}
	return Time(t), nil
}

func ParseDate(dateString string) (Date, error) {
	t, err := time.Parse(dateFormat, dateString)
	if err != nil {
		return Date(time.Time{}), err
	}
	return Date(t), nil
}

type Time time.Time

func (t *Time) UnmarshalJSON(s []byte) error {
	var (
		year int
		mon int
		mday int
		hour int
		min int
		sec int
	)
	if len(s) <= 2 || s[0] != '"' || s[len(s) - 1] != '"' {
		return fmt.Errorf("invalid time: %s", s)
	}
	var str = string(s[1:len(s) - 1])
	if n, err := fmt.Sscanf(str, "%d-%02d-%02d %02d:%02d:%02d", &year, &mon, &mday, &hour, &min, &sec); err != nil {
		return fmt.Errorf("invalid string(%s): %s", err.Error(), s)
	} else if n != 6 {
		return fmt.Errorf("invalid time: %s", s)
	}
	*t = Time(time.Date(year, time.Month(mon), mday, hour, min, sec, 0, time.Local))
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(timeFormat))
	return []byte(stamp), nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}


type Date time.Time

func (t *Date) UnmarshalJSON(s []byte) error {
	var (
		year int
		mon int
		mday int
	)
	if len(s) <= 2 || s[0] != '"' || s[len(s) - 1] != '"' {
		return fmt.Errorf("invalid date: %s", s)
	}
	var str = string(s[1:len(s) - 1])
	if n, err := fmt.Sscanf(str, "%d-%02d-%02d", &year, &mon, &mday); err != nil {
		return fmt.Errorf("invalid string(%s): %s", err.Error(), s)
	} else if n != 3 {
		return fmt.Errorf("invalid date: %s", s)
	}
	*t = Date(time.Date(year, time.Month(mon), mday, 0, 0, 0, 0, time.Local))
	return nil
}

func (t Date) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(dateFormat))
	return []byte(stamp), nil
}

func (t Date) String() string {
	return time.Time(t).Format(dateFormat)
}