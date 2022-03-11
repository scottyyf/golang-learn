package main

import "fmt"

type account struct {
	No string
	Pwd string
	Num float32
}

func (a account) Pay()  {
	fmt.Println(a.Num)
}

func (a account) SetNum()  {
	a.Num = 99.99
}

func NewFactory(no, pwd string, num float32) *account{
	account := &account{no, pwd, num}
	return account
}

func main() {
	factory := NewFactory("01", "88888888", 88.88)
	(*factory).Pay()
}
