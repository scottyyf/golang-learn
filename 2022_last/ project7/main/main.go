package main

import "fmt"
//import "sort"

type Monkey struct{
	Name string
}

func (m *Monkey) Climbing(){
	fmt.Println("this can climb", m.Name)
}

type LittleMonkey struct {
	Monkey
}

type Interface interface{
	Fly()
	Swim()
}

func (this LittleMonkey) Fly(){
	fmt.Println("little monkey fly")
}

func (this LittleMonkey) Swim(){
	fmt.Println("Little monkey swim", this.Name)
}


func main(){
	lt := LittleMonkey{Monkey{"tom"}}
	lt.Climbing()
	lt.Fly()
	lt.Swim()
}
