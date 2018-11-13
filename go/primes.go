package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}

	remainder := n % 6

	if remainder != 1 && remainder != 5 {
		return false
	}

	var sqrt_n float64
	tmp := float64(n)
	sqrt_n = math.Sqrt(tmp)

	for i := 5; i < int(sqrt_n); i += 6 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func Sqrt(x float64) float64 {
	z := float64(1)
	tmp := float64(0)
	for math.Abs(tmp-z) > 0.0000000001 {
		tmp = z
		z = (z + x/z) / 2
	}
	return z
}

func getPrimes(n int) {
	for i := 2; i <= n; i++ {
	}
}

func main() {
	fmt.Println(isPrime(61))
}
