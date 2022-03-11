package main

import "fmt"

var x = func(num1, num2 int) int{
	fmt.Println("r func")
	return num1 + num2
}(1, 2)



func main() {
	 var x = func(num1, num2 int) {
	 	fmt.Println("x")
	 }
	 x(1,2)
}
