package main

import "fmt"

const boilingF = 212.0

var pp = f() //p指针指向v的地址

func main() {
	var test int = 1
	var p = &test // p指针
	*p = 2 // 指针指向的内容变化2
	f := boilingF
	c := (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)
	fmt.Println(test)
	fmt.Println(*pp)
}


func f() *int  { //返回值是int类型的指针
	v := 1
	return &v
}
