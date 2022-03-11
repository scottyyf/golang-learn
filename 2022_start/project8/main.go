package main

import "fmt"

func main() {
	var i, j = -2, 2
	/*
	-2:
	原码: 1000 0010
	反码: 1111 1101
	补码: 1111 1110
	2:
	补码: 0000 0010
	异或:
	补码: 1111 1100
	反码: 1111 1011
	原码: 1000 0100
	      -4
	 */
	fmt.Println(i&j)
	fmt.Println(i|j)
	fmt.Println(i^j) // -4
	fmt.Println(1>>2)
	fmt.Println(2>>2)
	fmt.Println(4>>2)
	fmt.Println(8>>2)
	fmt.Println(-12345<<2)
}