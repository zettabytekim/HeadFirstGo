package main

import (
	"fmt"

	"github.com/zettabytekim/magazine"
)

func main() {
	Subscriber := magazine.Subscriber{Name: "Zetta Kim", Rate: 4.99, Active: true}
	fmt.Println("Name:", Subscriber.Name)
	fmt.Println("Rate:", Subscriber.Rate)
	fmt.Println("Active:", Subscriber.Active)
	fmt.Println()

	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
	fmt.Println()

	var address magazine.Address
	address.Street = "123 Oak St"
	address.City = "Omaha"
	address.State = "NE"
	address.PostalCode = "681111"
	fmt.Println(address)
	fmt.Println()

	address2 := magazine.Address{Street: "124 Oks St",
		City: "Omaha", State: "NE", PostalCode: "691111"}
	subscriber2 := magazine.Subscriber{Name: "Adam Signh"}
	subscriber2.HomeAddress = address2
	fmt.Println(subscriber2.HomeAddress)
}
