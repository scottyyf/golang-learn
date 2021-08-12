package main

import "fmt"

func newInt() *int {
	return new(int)
	// 等价下面这个
	//	var dummy int
	//	return &dummy
}

func main() {
	p := new(int)
	fmt.Println(*p)
}
