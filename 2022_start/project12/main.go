package main

import "fmt"

func looPrint() {
	var num = 0
	for {
		if num == 10 {
			break
		}
		fmt.Println("hello world")
		num++
	}
	fmt.Println(num)
}

func main() {
	looPrint()
}
