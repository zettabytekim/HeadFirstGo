package main

import (
	"HeadFirstGo/ch10/geo"
	"fmt"
)

func main() {
	coordinates := geo.Coordinates{}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)
}
