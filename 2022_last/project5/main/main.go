package main

import "fmt"

type AnInterFace interface {
	Say()
}

type BInterFace interface {
	Hello()
}

type EmInterFace interface {

}

type Student struct {
	name string
}

func (s Student) Say()  {
	fmt.Println("i am ", s.name)
}

func (s Student) Hello(){
	fmt.Println("Hello in student")
}

type Name struct {
	Names string
}

func (name *Name)Working(stu Student)  {
	fmt.Println(name.Names)
	stu.Say()
}

type Interger int

func main(){
	stu := Student{"scott"}
	x := Name{Names: "yg"}
	x.Working(stu)

	var b AnInterFace = stu
	b.Say()

	var c BInterFace = stu
	c.Hello()

	var d EmInterFace = stu
	fmt.Println(d)

	var integer = Interger(0)
	var e EmInterFace = integer
	fmt.Println(e)
}
