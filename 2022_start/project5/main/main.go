package main

import "fmt"

func main1(fahrenheit float32) float32 {
	var centiGrade float32
	centiGrade = 5.0 / 9 * (fahrenheit - 100)
	return centiGrade
}

func main() {
	var centiGrade = main1(134.2)
	fmt.Println(centiGrade)
}
