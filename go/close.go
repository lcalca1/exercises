package main

import "fmt"
import "time"

func main() {
	go func() {
		time.Sleep(1 * time.Hour)
	}()

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//close(c)
	}()

	for i := range c {
		fmt.Printf("%d", i)
	}
	fmt.Println("Finished")
}
