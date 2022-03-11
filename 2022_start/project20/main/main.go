package main

import (
	"fmt"
	utils "golang-learn/2022_start/project20/uitls"
)

var name = sum(1, 2)

func init() {
	fmt.Println("main init called")
}


func sum(n1, n2 float32) float32{
	fmt.Println("var")
	return n1 + n2
}

func swapValue(n1, n2 *int){
	tmp := *n1
	*n1 = *n2
	*n2 = tmp
	//*n1, *n2 = *n2, *n1
}

func main() {
	fmt.Println(sum(1, 2.99))
	var n1, n2 = 1, 2
	swapValue(&n1, &n2)
	fmt.Println(n1, " ",n2)
	fmt.Println(utils.Age, utils.Name)
}