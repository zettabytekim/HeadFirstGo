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
	fmt.Println()

	event := calendar.Event{}
	err = event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(event.Title)
	err = event.SetYear(2020)
	if err != nil {
		log.Fatal()
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal()
	}
	err = event.SetDay(16)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
	fmt.Println()

	fmt.Println(event.Date.Year())
	fmt.Println(event.Date.Month())
	fmt.Println(event.Date.Day())
	fmt.Println()
}
