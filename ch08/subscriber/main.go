package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

type myStruct struct {
	myField int
}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly Rate:", s.rate)
	fmt.Println("Active:", s.active)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func double(number *int) {
	*number *= 2
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Zetta Kim")
	subscriber1.rate = 4.99
	printInfo(subscriber1)
	fmt.Println()

	subscriber2 := defaultSubscriber("Brian Adams")
	printInfo(subscriber2)
	fmt.Println()

	amount := 6
	double(&amount)
	fmt.Println(amount)
	fmt.Println()

	var value int = 2
	var pointer *int = &value
	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Println()

	var value2 myStruct
	value2.myField = 3
	var pointer2 *myStruct = &value2
	fmt.Println(pointer2.myField)
	fmt.Println()

	var value3 myStruct
	var pointer3 *myStruct = &value3
	pointer3.myField = 9
	fmt.Println(pointer3.myField)
	fmt.Println()

	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)
	fmt.Println()

}
