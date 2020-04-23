package main

import (
	"HeadFirstGo/ch04/dates"
	"fmt"
)

func main() {
	days := 3
	fmt.Println("You appointment is in", days, "days")
	fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days")
}
