package main

import "fmt"

var v = 1

func main()  {
	incr(&v)
	fmt.Println(v)
}

func incr(p *int) int  {//参数p是int类型的指针，返回值是int类型
	*p += 10 //p指向的值+1
	return *p
}
