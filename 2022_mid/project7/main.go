package main

import (
	"fmt"
)

func sliceDemo() {
	var arr = [...]int{1, 2, 3, 4, 5}
	var slice = arr[1:3]

	fmt.Println(&arr[1])
	fmt.Printf("%p", slice)
	fmt.Println()

	slice[0] = 100
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
}

func testDemo() {
	var slice = make([]int, 4, 10)
	for i := 0; i < len(slice); i++ {
		//slice[i] = i
	}
	fmt.Println(len(slice))
	fmt.Println(slice)
	//slice[7] = 0
	fmt.Println(slice)
}

func testSlice() {
	var arr = [...]int{10, 20, 30, 40, 50}
	slice := arr[:]
	slice2 := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	slice = append(slice, slice2...)
	fmt.Println(slice)
	fmt.Println(slice2)
	sliceW := append([]byte("hello "), "world"...)
	for i := 0; i < len(sliceW); i++ {
		fmt.Printf("%c", sliceW[i])
	}
	fmt.Println(sliceW)
	var sliceM = make([]byte, 10)
	copy(sliceM, sliceW)
	for i := 0; i < len(sliceM); i++ {
		fmt.Printf("%c", sliceM[i])
	}
	fmt.Println(string(sliceW))
}

func stringTest(){
	str := "19239716"
	tmp:=[]rune(str)
	tmp[0] = '中'
	ret := string(tmp)
	fmt.Println(ret)
}

func main() {
	//sliceDemo()
	//testDemo()
	//testSlice()
	//stringTest()
	var x = []byte("hello")
	x = append(x, "world"...)
	fmt.Println(string(x))
	//fmt.Println(strings.Join(x, "world"))
}
