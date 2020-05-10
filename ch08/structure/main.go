package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func minumumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)
	fmt.Println()

	myStruct.number = 3.14
	fmt.Println(myStruct.number)
	fmt.Println()

	var subscriber1 subscriber
	subscriber1.name = "Adam Signh"
	subscriber1.rate = 4.99
	subscriber1.active = true
	fmt.Println("Name:", subscriber1.name)
	fmt.Println("Monthly Rate:", subscriber1.rate)
	fmt.Println("Active:", subscriber1.active)

	var subscriber2 subscriber
	subscriber2.name = "Beth Ryan"
	subscriber2.rate = 5.12
	subscriber2.active = false
	fmt.Println("Name:", subscriber2.name)
	fmt.Println("Monthly Rate:", subscriber2.rate)
	fmt.Println("Active:", subscriber2.active)
	fmt.Println()

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts)
	fmt.Println()

	p := minumumOrder("Hex bolts")
	fmt.Println(p.description, p.count)

}
