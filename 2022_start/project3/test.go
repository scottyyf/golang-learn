package main

import (
	"fmt"
	"strconv"
)

func anal(){
	var str = "0"
	b, e := strconv.ParseBool(str)
	fmt.Println(b, e)
	fmt.Printf("%T",b)
	fmt.Println()
	fmt.Printf("%T",str)

	var str2 = "12345611111"
	i, _ := strconv.ParseInt(str2, 10, 0)
	fmt.Println()
	fmt.Println(i)
	fmt.Printf("%T", i)
}

func main() {
	anal()
}