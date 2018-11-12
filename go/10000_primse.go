package main

import "fmt"

func Processor(seq chan int, wait chan struct{}) {
	go func() { // 把当前 seq 第一个数取出并输出（这个数是素数），然后剩下的逐个除以之前取出的素数，过滤掉可以除尽的，然后再放入 seq 中
		prime, ok := <-seq

		if !ok {
			close(wait)
			return
		}

		fmt.Println("[", prime, "]")
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}

func main() {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < 10000000; num++ {
		origin <- num
	}
	close(origin)
	<-wait
}
