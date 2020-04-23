package main

import (
	// "HeadFirstGo/ch04/keyboard"
	"fmt"
	"log"

	"github.com/zettabytekim/HeadFirstGo/ch04/keyboard"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}
