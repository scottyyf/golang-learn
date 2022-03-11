package service

import (
	"fmt"
	"golang-learn/2022_last/project10/customerManage/model"
)

type CustomerService struct {
	// 完成对view的各种操作逻辑实现
	customers   []*model.Customer
	customerNum int
}

func (c *CustomerService) GetCustomerNum() int {
	return c.customerNum
}

func (c *CustomerService) List() []*model.Customer {
	return c.customers
}

func (c *CustomerService) Add(customer *model.Customer) bool {
	// 加一个id
	c.customerNum++
	customer.Id = c.customerNum
	c.customers = append(c.customers, customer)
	return true
}

func (c *CustomerService) Del(id int) bool {
	index := c.FindById(id)
	if index != -1 {
		c.customers = append(c.customers[:index],
			c.customers[index+1:]...)
		return true
	}
	return false
}

func (c *CustomerService) Modify(index int,name string,
	gender string, email string, tel string, age int) bool {
	oldCustomer := c.customers[index]

	fmt.Println(name, "0000000000")
	if name == ""{
		name = oldCustomer.Name
	}

	if gender == ""{
		gender = oldCustomer.Gender
	}

	if email == ""{
		email = oldCustomer.Email
	}

	if tel == ""{
		tel = oldCustomer.Phone
	}

	if age == 0{
		age = oldCustomer.Age
	}

	customer := &model.Customer{
		Name: name,
		Gender: gender,
		Email: email,
		Phone: tel,
		Age: age,
	}
	oldCustomer.Update(customer)
	return true
}

func (c *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(c.customers); i++ {
		if c.customers[i].Id == id {
			return i
		}
	}
	return index
}

func NewCustomerService() *CustomerService {
	//
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(
		"张三", "男",
		18, "12345", "email.qq.com",
	)
	customer.Id = 1
	customerService.customers = append(customerService.customers,
		customer)

	return customerService
}
