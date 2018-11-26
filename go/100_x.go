package main

import (
	"fmt"
	"math/rand"
	"time"
)

func do_stuff(x int) int { //一个比较耗时的事情
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	return 100 - x
}

func branch(x int) chan int {
	ch := make(chan int)
	go func() {
		ch <- do_stuff(x)
	}()

	return ch
}

func fanIn1(chs ...chan int) chan int {
	ch := make(chan int)

	for _, c := range chs {
		// 注意此处明确传值
		go func(c chan int) { ch <- <-c }(c) //复合
	}

	return ch
}

func fanIn(chs ...chan int) chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < len(chs); i++ {
			select {
			case v := <-chs[i]:
				ch <- v
			}
		}
	}()

	return ch
}

func main() {
	result := fanIn(branch(1), branch(2), branch(3), branch(4))

	for i := 0; i < 4; i++ {
		fmt.Println(<-result)
	}
}
