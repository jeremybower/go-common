package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Date struct {
	Valid bool
	Year  int
	Month time.Month
	Day   int
}

func NewDate(year int, month time.Month, day int) Date {
	d := Date{}
	d.assign(year, month, day)
	return d
}

func (d *Date) assign(year int, month time.Month, day int) {
	d.Year = year
	d.Month = month
	d.Day = day

	t := time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
	d.Valid = t.Year() == year && t.Month() == time.Month(month) && t.Day() == day
}

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

var ErrInvalidDate = errors.New("invalid date")

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidDate, err)
	}

	parts := strings.Split(s, "-")
	if len(parts) != 3 {
		return fmt.Errorf("%w: %s", ErrInvalidDate, s)
	}

	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidDate, s)
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidDate, s)
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidDate, s)
	}

	d.assign(year, time.Month(month), day)
	return nil
}
