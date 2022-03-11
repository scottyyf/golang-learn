package main

import (
	"golang-learn/2022_last/project10/customerManage/service"
	"golang-learn/2022_last/project10/customerManage/view"
)

func main() {
	var customerView = &view.CustomerView{
		Key: "",
		ExitTag: false,
	}

	customerView.Service = service.NewCustomerService()
	customerView.MainMenu()
}
