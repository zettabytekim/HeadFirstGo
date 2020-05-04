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
}
