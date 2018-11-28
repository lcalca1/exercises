package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}
func m() {
	a = "D"  //global scope
	a := "O" //local scope
	a = "D"  //local scope
	print(a)
}
