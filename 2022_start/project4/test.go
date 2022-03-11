package main

import "fmt"

func pointer(i *int) *int {
	var ptr = i
	fmt.Println("ptr content is memory location", ptr)
	fmt.Println("ptr content's content is ", *ptr)
	return ptr
}


func main() {
	var i = 1234
	//var ptr *int = &i
	ptr := pointer(&i)
	fmt.Println("memory location is ", ptr)
	*ptr = 20
	fmt.Println(i)
}
