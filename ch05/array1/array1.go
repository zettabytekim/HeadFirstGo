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
	fmt.Println(primes[1])
	fmt.Println(primes[2])
	fmt.Println(primes[3])

	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	fmt.Println(dates[0])

	var counters [3]int
	counters[0]++
	counters[0]++
	counters[2]++
	fmt.Println(counters[0], counters[1], counters[2])

	var notes2 [7]string = [7]string{"도", "레", "미", "파", "솔", "라", "시"}
	fmt.Println(notes2[3], notes2[6], notes2[0])

	var primes2 [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes2[0], primes2[2], primes2[4])

	fmt.Println(notes2)
	fmt.Println(primes2)

	fmt.Printf("%#v\n", notes2)
	fmt.Printf("%#v\n", primes2)

	for i := 0; i < 7; i++ {
		fmt.Println(i, notes2[i])
	}

	// for i := 0; i <= 7; i++ {
	// 	fmt.Println(i, notes2[i])
	// }

	fmt.Println(len(notes2))
	fmt.Println(len(primes2))

	for i := 0; i < len(notes2); i++ {
		fmt.Println(i, notes2[i])
	}

	fmt.Println("\nfor...range")
	for index, notes := range notes2 {
		fmt.Println(index, notes)
	}
}
