package main

import "fmt"

var complete chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	complete <- 0
}

func main() {
	var messages chan string = make(chan string)

	go loop()
	go func(message string) {
		messages <- message
	}("Ping!")

	fmt.Printf(<-messages)
	fmt.Println()
	fmt.Printf("%d", <-complete)
}
