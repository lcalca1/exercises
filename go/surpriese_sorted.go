// 一种令人意想不到的排序方法
package main

import (
	"fmt"
	"time"
)

func main() {
	var arr = [...]int{300, 60, 1, 9, 4, 2, 7, 5}
	//var sorted, wait = make(chan int), make(chan struct{})
	var sorted = make(chan int)

	for i := 0; i < len(arr); i++ {
		go func(i int) {
			if arr[i] > 0 {
				time.Sleep(time.Duration(arr[i]) * time.Millisecond)
			}
			sorted <- arr[i]
		}(i)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Print(<-sorted, " ")
	}
}
