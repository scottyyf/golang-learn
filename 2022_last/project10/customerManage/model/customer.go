package model

import "fmt"

type Customer struct {
	// 表示一个客户信息
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) *Customer {
	// new 一个customer实例
	return &Customer{
		Id: id,
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}

func (c *Customer) GetInfo() string {
	info := fmt.Sprintf("%d\t%v\t%v\t%v\t%v\t%v\n",
		c.Id, c.Name, c.Gender, c.Age,
		c.Phone, c.Email)
	return info
}