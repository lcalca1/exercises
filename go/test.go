package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	a := 3
	for i := 0; i < 5; i++ {
		fmt.Println(a)
		a := 1
		a++
		if i == 4 {
			fmt.Println(a)
		}
	}
	fmt.Println(a)
}
