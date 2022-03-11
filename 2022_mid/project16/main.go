package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) speak() {
	fmt.Println(p.Name, "is a goodman")
}

func (p Person) jisuan()  {
	ret := 0
	for i := 1; i < 1001; i++ {
		ret += i
	}
	fmt.Println(ret)
}

func (p Person) jisuan2(num int)  {
	ret := 0
	for i := 1; i <= num; i++ {
		ret += i
	}
	fmt.Println(ret)
}

func (p Person) getSum(num1, num2 int) int{
	ret := num1 + num2
	return ret
}

type Circle struct {
	Radis float64
}

type integer int

func (i integer) print() integer{
	fmt.Println("i=", i)
	return i
}

func (c *Circle) area() float64 {
	fmt.Printf("%p", c)
	fmt.Println()
	return 3.14 * (*c).Radis * (*c).Radis
}

func (c *Circle) String() string{
	str := fmt.Sprintf("%v ^^^^", c.Radis)
	return str
}

func main() {
	var c = Circle{2.91}
	fmt.Println(&c)
}
