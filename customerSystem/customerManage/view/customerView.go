package view

import (
	"fmt"
	"golang-learn/2022_last/project10/customerManage/model"
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
	fmt.Println("编号\t\t姓名\t\t性别\t\t年龄\t\t电话\t\t邮箱")
	for _, value := range ret {
		fmt.Println(value.GetInfo())
	}
	fmt.Println("----------客户列表完成----------")
}

func (c *CustomerView) Add()  {
	var name, gender, email, tel string
	var age int
	fmt.Println("----------添加客户----------")
	fmt.Printf("姓名:\t")
	fmt.Scanln(&name)
	fmt.Printf("性别:\t")
	fmt.Scanln(&gender)
	fmt.Printf("邮箱:\t")
	fmt.Scanln(&email)
	fmt.Printf("电话:\t")
	fmt.Scanln(&tel)
	fmt.Printf("年龄:\t")
	fmt.Scanln(&age)

	customer := model.NewCustomer(
		name, gender, age, tel, email)
	ret := c.Service.Add(customer)
	if ret{
		fmt.Println("添加用户完成")
	}else {
		fmt.Println("添加用户失败")
	}
}

func (c *CustomerView) Del()  {
	fmt.Println("----------删除客户--------")
	var key int
	fmt.Printf("请选择删除客户的编号(0退出):\t")
	fmt.Scanln(&key)
	if c.Service.Del(key){
		fmt.Println("删除成功")
	}else {
		fmt.Println("删除失败")
	}
}

func (c *CustomerView) Modify()  {
	fmt.Println("----------修改客户----------")
	var key int
	fmt.Println("请选择修改客户的编号(0退出):\t")
	fmt.Scanln(&key)
	index := c.Service.FindById(key)
	if index == -1{
		fmt.Println("找不到这个客户编号")
		return
	}
	var name, gender, email, tel string
	var age int

	fmt.Printf("姓名:\t")
	fmt.Scanln(&name)
	fmt.Printf("性别:\t")
	fmt.Scanln(&gender)
	fmt.Printf("邮箱:\t")
	fmt.Scanln(&email)
	fmt.Printf("电话:\t")
	fmt.Scanln(&tel)
	fmt.Printf("年龄:\t")
	fmt.Scanln(&age)
	c.Service.Modify(index, name, gender, email, tel, age)
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
		fmt.Println("\t请选择(1-5): ")
		labels:
		fmt.Scanln(&c.Key)
		switch c.Key {
		case "1":
			c.Add()
		case "2":
			c.Modify()
		case "3":
			c.Del()
		case "4":
			c.List()
		case "5":
			c.ExitTag = true
		case "":
			goto labels
		default:
			fmt.Println("invalid input, pls input again")
		}

		if c.ExitTag{
			break
		}
		c.Key = ""
	}

	fmt.Println("你退出了客户管理系统")
}
