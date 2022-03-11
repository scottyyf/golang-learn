package main

import "fmt"

func chickenInfo(){
	//var n = [...]int{1,23}
	//var name [4]float64 = [...]float64{1,2,3, 9}
	//name[100] = 1.0
	//fmt.Printf("%T", name)
	var nm = [3]string{1:"tom", 2:"moniso", 0:"yakebu"}
	//fmt.Println(nm, n)
	for index, value := range nm{
		fmt.Println(index, value)
	}
	// sliceX list type
	var slice [3]*float64
	fmt.Println(slice)
}

func test(arr *[3]int){
	(*arr)[0] = 99
	fmt.Println(*arr)
}

func main() {
	chickenInfo()
	arr := [...]int{11,22,33}
	test(&arr)
	fmt.Println(arr)
}
