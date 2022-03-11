package main

import (
	"fmt"
	"golang-learn/2022_last/project2/model"
)

func main() {
	person := model12.NewPerson("scott")
	person.SetAge(0)
	fmt.Println(person)
}
