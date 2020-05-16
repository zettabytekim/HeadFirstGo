package main

import (
	"HeadFirstGo/ch10/calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}
