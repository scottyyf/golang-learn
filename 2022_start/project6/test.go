package main

import "fmt"

func main() {
	var a, b = 100, 200
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}