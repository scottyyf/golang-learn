package main

import "fmt"

func classScore() (float64, int){
	var ret float64 = 0
	var passedNum int = 0
	fmt.Println("input class score")
	for student := 0; student < 5; student++ {
		var score float64
		fmt.Scan(&score)
		if score >= 60{
			passedNum += 1
		}
		ret += score
	}
	ret = ret / 5.0
	fmt.Println("ava score is ", ret)
	fmt.Println("class passed ", passedNum)
	return ret, passedNum
}

func schoolScore(){
	var ret float64 = 0
	var passed int = 0
	for i := 0; i < 5; i++ {
		result, passedNum := classScore()
		ret += result
		passed += passedNum
	}
	ret = ret / 5.0 / 5.0
	fmt.Printf("ava school score is %.2f", ret)
	fmt.Println()
	fmt.Println("passed ", passed)
}

func main() {
	schoolScore()
}
