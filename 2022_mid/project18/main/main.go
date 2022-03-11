package main

import (
	"fmt"
	"golang-learn/2022_mid/project18/factory"
)

func main() {
	var stu = factory.NewStudent("x", 89.0)
	fmt.Println(stu.Name, stu.GetScore())
}
