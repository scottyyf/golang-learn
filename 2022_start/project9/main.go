package main

import "fmt"

func compareAndPrint(x int8, y int8) {
	if x+y >= 50 {
		fmt.Println("hello world")
	} else {
		fmt.Println("less than 50")
	}
}

func compareAndPrint2(x float64, y float64) {
	if x > 10.0 && y < 20.0 {
		fmt.Println(x + y)
	} else {
		fmt.Println("compareAndPrint2 to nothing")
	}
}

func compareAndPrint3(x int32, y int32) {
	var ret = x + y
	if ret%3 == 0 && ret%5 == 0 {
		fmt.Printf("result of %d plus %d can divide by 3 and 5", x, y)
	} else {
		fmt.Printf("result of %d plus %d can not divide by 3 and 5",
			x, y)
	}
	fmt.Println()
}

func isPrime(year int){
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0{
		fmt.Println("yes")
	}else {
		fmt.Println("no")
	}
}

func main() {
	compareAndPrint(32, 20)
	compareAndPrint2(32.0, 10)
	compareAndPrint3(30, 60)
	isPrime(2019)
}
