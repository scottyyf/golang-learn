package view

import (
	"fmt"
	"golang-learn/2022_last/project10/customerManage/service"
)

type CustomerView struct {
	// 接受输入
	Key string
	ExitTag bool
	Service *service.CustomerService
}

func (c *CustomerView) List()  {
	ret := c.Service.List()
	fmt.Println("----------客户列表----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, value := range ret {
		fmt.Println(value.GetInfo())
	}
	fmt.Println("----------客户列表完成----------")
}

func (c *CustomerView) MainMenu()  {
	for {
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println()
		fmt.Println("\t1.添加客户")
		fmt.Println("\t2.修改客户")
		fmt.Println("\t3.删除客户")
		fmt.Println("\t4.客户列表")
		fmt.Println("\t5.退出")
		fmt.Println("\t请选择(1-5: ")

		fmt.Scanln(&c.Key)
		switch c.Key {
		case "1":
			//
		case "2":
			//
		case "3":
			//
		case "4":
			c.List()
		case "5":
			//
			c.ExitTag = true
		default:
			fmt.Println("invalid input, pls input again")
		}

		if c.ExitTag{
			break
		}
	}

	fmt.Println("你退出了客户管理系统")
}
