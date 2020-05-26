package main

import (
	"fmt"
	"log"
)

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheatung by %0.2f degrees!", o)
}

func chechTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func main() {
	var err error = chechTemperature(121.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}
