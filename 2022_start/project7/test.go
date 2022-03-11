package main

import "fmt"

func main() {
	var (
		name     string
		age      int
		salary   float32
		isPassed bool
	)
	fmt.Scan(&name, &age, &salary, &isPassed)

	//fmt.Scanln(&name)
	//fmt.Scanln(&age)
	//fmt.Scanln(&salary)
	//fmt.Scanln(&isPassed)

	//fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPassed)

	fmt.Println(name, age, salary, isPassed)
}
