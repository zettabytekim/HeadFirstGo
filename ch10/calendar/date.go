package calendar

import "errors"

// Date struct
type Date struct {
	year  int
	month int
	day   int
}

// SetYear Method
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// Year Method
func (d *Date) Year() int {
	return d.year
}

// Month Method
func (d *Date) Month() int {
	return d.month
}

// Day Method
func (d *Date) Day() int {
	return d.day
}

// SetMonth Method
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// SetDay Method
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
