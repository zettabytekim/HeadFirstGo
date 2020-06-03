package main

import "fmt"

func greeting(myChannel chan string) {
	myChannel <- "Hi"
}
func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<-myChannel)
}
