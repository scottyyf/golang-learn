package main

import (
	"fmt"
	"strconv"
)

type student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (stu student) say() string {
	stu.id = 90
	return stu.name + ";\t" + stu.gender + ";\t" + strconv.Itoa(stu.age) + ";\t" +
		strconv.Itoa(stu.id) + ";\t" +
		strconv.FormatFloat(stu.score, 'f', 2, 64)
}

func main() {
	var xiaoMing = student{
		name: "yes",
		age: 8,
	}
	fmt.Printf("%s", xiaoMing.say())
	fmt.Println(xiaoMing)
}
