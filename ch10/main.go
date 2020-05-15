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

	err = date.SetMonth(13)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(date)
	// date.year = 2019
	// date.month = 14
	// date.day = 33
	// fmt.Println(date)

}
