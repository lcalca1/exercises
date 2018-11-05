package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	timeout := time.After(time.Second * 2)
	t1 := time.NewTimer(time.Second * 3)
	var i int
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("channel sign")
				return
			case <-t1.C:
				fmt.Println("3s 定时任务")
			case <-timeout:
				i++
				fmt.Println("2s 定时任务")
			case <-time.After(time.Second * 2):
				fmt.Println("4s timeout..")
			default:
				fmt.Println("default")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	time.Sleep(time.Second * 9)
	close(c)
	time.Sleep(time.Second * 2)
	fmt.Println("exit....")
}
