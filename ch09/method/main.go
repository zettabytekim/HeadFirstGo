package main

import "fmt"

type MyType string

type Number int

func (n *Number) Double() {
	*n *= 2
}

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}
func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer receiver")
}

func main() {
	value := MyType("a MyType value")
	value.sayHi()
	anotherValue := MyType("another value")
	anotherValue.sayHi()
	fmt.Println()

	number := Number(4)
	fmt.Println("Original:", number)
	number.Double()
	fmt.Println("after:", number)

	value2 := MyType("a value")
	pointer2 := &value2
	value2.method()
	value2.pointerMethod()
	pointer2.method()
	pointer2.pointerMethod()
}
