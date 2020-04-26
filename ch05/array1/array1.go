package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	notes[0] = "도"
	notes[1] = "레"
	notes[2] = "미"
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	fmt.Println(dates[0])
}
