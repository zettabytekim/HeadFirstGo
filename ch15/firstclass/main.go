package main

import "fmt"

func SayHi() {
	fmt.Println("Hi")
}
func SayBye() {
	fmt.Println("Bye")
}
func twice(theFunction func()) {
	theFunction()
	theFunction()
}
func main() {
	twice(SayHi)
	twice(SayBye)
}
