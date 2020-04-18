package main

import (
	"fmt"
)

// func paintNeeded(width float64, height float64) {
// 	area := width * height
// 	fmt.Printf("%.2f liters needed\n", area/10.0)
// }
var metersPerLiter float64

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / metersPerLiter
}

func main() {
	// paintNeeded(4.2, 3.0)
	// paintNeeded(5.2, 3.5)
	// paintNeeded(5.0, 3.3)

	metersPerLiter = 10.0
	fmt.Printf("%.2f", paintNeeded(4.2, 3.0))
}
