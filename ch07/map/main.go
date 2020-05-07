package main

import (
	"fmt"
)

func main() {
	// var ranks map[string]int
	// ranks = make(map[string]int)

	// ranks["gold"] = 1
	// ranks["silver"] = 2
	// ranks["bronze"] = 3

	ranks := map[string]int{"gold": 1, "silver": 2, "bronze": 3}

	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["silver"])
	fmt.Println(ranks["gold"])
	fmt.Println()

	// elements := make(map[string]string)
	// elements["H"] = "Hydrogen"
	// elements["Li"] = "Lithum"

	elements := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithum",
	}

	fmt.Println(elements["H"])
	fmt.Println(elements["Li"])
	fmt.Println()

	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 79.99
	clothing := map[string]float64{"pants": 59.99, "shirts": 39.99}
	fmt.Println("Earrings:", jewelry["earrings"])
	fmt.Println("Necklace:", jewelry["necklace"])
	fmt.Println("Shirts:", clothing["shirts"])
	fmt.Println("Pants:", clothing["pants"])

	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
	fmt.Println()

	_, ok = counters["b"]
	fmt.Println(ok)
	_, ok = counters["c"]
	fmt.Println(ok)
	fmt.Println()

	status("Alma")
	status("Carl")
	fmt.Println()
	wordCount()

	fmt.Println()
	mapDelete()
}

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Roit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}

func wordCount() {
	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
	}
	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}
	}
}

func mapDelete() {
	var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)

	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	fmt.Println()

	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
}
