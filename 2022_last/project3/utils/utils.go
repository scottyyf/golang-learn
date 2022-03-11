package utils

import "fmt"

type C struct {

}

func (i *C) SayOk()  {
	fmt.Println("C sayok")
}

type A struct {
	Name string
	Age int
}

func (receiver *A) SayOk()  {
	fmt.Println("A sayOk", receiver.Name)
}

func (receiver *A) Hello()  {
	fmt.Println("A hello", receiver.Name)
}

func (receiver *A) Wataxi()  {
	fmt.Println("A wataxi", receiver.Name)
}

//func (i *B) SayOk()  {
//	fmt.Println("B sayok", i.Name)
//}

type B struct {
	X A
	C
	int
	Name string
}

func (i *B) Setint(num int)  {
	i.int = num
}
