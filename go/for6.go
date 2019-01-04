package main

import "fmt"

func main() {
LABEL1:
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 2; j++ {
			if j == 1 {
				continue LABEL1
			}
			fmt.Printf("i is %d, and j is %d\n", i, j)
		}
	}
}
