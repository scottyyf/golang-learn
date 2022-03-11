package service

import "golang-learn/2022_last/project10/customerManage/model"

type CustomerService struct {
	// 完成对view的各种操作逻辑实现
	customers []*model.Customer
	customerNum int
}

func (c *CustomerService) List() []*model.Customer {
	return c.customers
}

func NewCustomerService() *CustomerService {
	//
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(
		1, "张三","男",
		18,"12345", "email.qq.com",
		)
	customerService.customers = append(customerService.customers,
		customer)

	return customerService
}
