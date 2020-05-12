package main

import (
	"fmt"

	"github.com/zettabytekim/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Zetta Kim"}
	// subscriber.HomeAddress.Street = "123 Oak St"
	// subscriber.HomeAddress.City = "Omaha"
	// subscriber.HomeAddress.Street = "NE"
	// subscriber.HomeAddress.PostalCode = "68111"
	// 익명필드 사용
	// subscriber.Address.Street = "123 Oak St"
	// subscriber.Address.City = "Omaha"
	// subscriber.Address.Street = "NE"
	// subscriber.Address.PostalCode = "68111"
	//구조체 임베딩에 의한 승격
	subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"
	subscriber.Street = "NE"
	subscriber.PostalCode = "68111"

	// fmt.Println("Subscriber Name:", subscriber.Name)
	// fmt.Println("Street:", subscriber.HomeAddress.Street)
	// fmt.Println("City:", subscriber.HomeAddress.City)
	// fmt.Println("Postal Code:", subscriber.HomeAddress.PostalCode)

	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.Address.Street)
	fmt.Println("City:", subscriber.Address.City)
	fmt.Println("Postal Code:", subscriber.Address.PostalCode)
	fmt.Println()

	employeee := magazine.Employee{Name: "Mariah Carey"}
	employeee.Address.Street = "123333 Oak St"
	employeee.Address.City = "Omaha2"
	employeee.Address.Street = "NE2"
	employeee.Address.PostalCode = "78654"
	fmt.Println("Subscriber Name:", employeee.Name)
	fmt.Println("Street:", employeee.Address.Street)
	fmt.Println("City:", employeee.Address.City)
	fmt.Println("Postal Code:", employeee.Address.PostalCode)
	fmt.Println()
}
