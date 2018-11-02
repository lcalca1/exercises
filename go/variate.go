package main

import "fmt"

var a string = "菜鸟"
var b = "runoob.com"
var c bool
var (
	d int
	e bool
)

func main() {
	g, h := 123, "hello"
	println(a, b, g, h)
	a, b = b, a
	fmt.Println(a, b, g, h)

	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)

	// 旗标、宽度、精度、索引
	fmt.Printf("|%0+- #[1]*.[2]*[3]d|%0+- #[1]*.[2]*[4]d|\n", 8, 4, 32, 64)
	fmt.Printf("|%+- [1]*.[2]*[3]d|%0+- #[1]*.[2]*[4]d|\n", 3, 3, 32, 64)

	const (
		i = 1 << iota
		j
		k
		l
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

	var ptr *int
	var s int = 5
	ptr = &s
	s = 6
	ptr = &g
	println(s)
	println(*ptr)
}
