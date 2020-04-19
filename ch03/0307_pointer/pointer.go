package main

import "fmt"

func main() {
	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)

	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)

	myBool := true
	myBoolPinter := &myBool
	fmt.Println(myBoolPinter)
	fmt.Println(*myBoolPinter)

	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}
