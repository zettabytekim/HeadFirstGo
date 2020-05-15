package main

import "fmt"

// Date struct
type Date struct {
	Year  int
	Month int
	Day   int
}

// SetYear Method
func (d *Date) SetYear(year int) {
	d.Year = year
}

// SetMonth Method
func (d *Date) SetMonth(month int) {
	d.Month = month
}

// SetDay Method
func (d *Date) SetDay(day int) {
	d.Day = day
}

func main() {
	// date := Date{Year: 2020, Month: 5, Day: 16}
	// fmt.Println(date)	date := Date{}
	date := Date{}
	date.SetYear(2020)
	date.SetMonth(5)
	date.SetDay(15)
	fmt.Println(date)
}
