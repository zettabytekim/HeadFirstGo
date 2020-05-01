package main

import "fmt"

func main() {
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(len(primes))

	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}

	for _, letter := range letters {
		fmt.Println(letter)
	}

	notes := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	primes2 := []int{
		2,
		3,
		5,
	}
	fmt.Println(primes2[0], primes2[1], primes2[2])
	fmt.Println("-----------")

	numbers2 := make([]float64, 3)
	numbers2[0] = 19.7
	numbers2[2] = 25.2
	for i, number2 := range numbers2 {
		fmt.Println(i, number2)
	}
	var letters2 = []string{"a", "b", "c"}
	for i, letter2 := range letters2 {
		fmt.Println(i, letter2)
	}

	arr1 := [5]string{"a", "b", "c", "d", "e"}
	slic1 := arr1[0:3]
	arr1[1] = "X"
	fmt.Println(arr1)
	fmt.Println(slic1)

	arr2 := [5]string{"f", "g", "h", "i", "j"}
	slic2 := arr2[2:5]
	slic2[1] = "X"
	fmt.Println(arr2)
	fmt.Println(slic2)

	slic3 := []string{"a", "b"}
	fmt.Println(slic3, len(slic3))
	slic3 = append(slic3, "c")
	fmt.Println(slic3, len(slic3))
	slic3 = append(slic3, "d", "e")
	fmt.Println(slic3, len(slic3))

	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])

	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlcie: %#v, stringSlice: %#v\n", intSlice, stringSlice)
	fmt.Println(len(intSlice))

	var slic4 []string
	if len(slic4) == 0 {
		slic4 = append(slic4, "first item")
	}
	fmt.Printf("%#v\n", slic4)
}
