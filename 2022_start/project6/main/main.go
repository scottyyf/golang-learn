package main

import "fmt"

func isOldMan(age int, sex string) bool {
	if sex == "male" && age >= 65 {
		return true
	}

	if sex == "female" && age >= 60 {
		return true
	}

	return false
}

func isBeautiful(age int, sex string) bool {
	//男人是泥巴, 不说你丑就不错了
	if age < 415 || sex != "female"{
		return true
	}
	return false
}

func main() {
	var age = 64
	var sex = "female"
	fmt.Println("the person is old? ", isOldMan(age, sex))
	fmt.Println(isBeautiful(age, sex))
}
