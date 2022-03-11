package utils

import "fmt"

type customer struct {
	name string
	sex string
	age int
	tel string
	mail string

	key string
}



func NewCustomer(name string) *customer{
	return &customer{name: name}
}


func (c *customer) showChoice()  {
	fmt.Println("----------客户信息管理软件----------")
	fmt.Println("1.添加客户")
	fmt.Println("2.修改客户")
	fmt.Println("3.删除客户")
	fmt.Println("4.客户列表")
	fmt.Println("5.退出")
	fmt.Println("请选择(1-5: ")
}

func (c *customer) MainMenu()  {
	c.showChoice()
	for {
		fmt.Scanln(&c.key)
		switch c.key {
		case "1":
		case "2":
		case "3":

		case "4":
			c.showChoice()
		case "5":

		default:
			fmt.Println("choose a num again")
		}
	}
}
