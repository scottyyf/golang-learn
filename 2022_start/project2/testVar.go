package main

import "fmt"

var (
	nn = "xx"
	mm = 2
)

func test() {
	var n1, n2, n3, x = 1, 2, 3, "x"
	fmt.Println(n1, n2, n3, x)
	fmt.Println(nn, mm)
}

func main() {
	test()
}
