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
}
