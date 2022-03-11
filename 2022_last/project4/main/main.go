package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

type Camera struct {
	Name string
}

type Usb2 interface {
	Start()
	Stop()
}

func (p Phone) Start() {
	fmt.Println("phone start work", p.Name)
}

func (p Phone) Stop() {
	fmt.Println("phone stop work", p.Name)
}

func (p Camera) Start() {
	fmt.Println("camera start work", p.Name)
}

func (p Camera) Stop() {
	fmt.Println("camera stop work", p.Name)
}

type Computer struct {
}

func (receiver Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	var usbArr [3]Usb

	usbArr[0] = Phone{"huawei"}
	usbArr[1] = Phone{"xiaomi"}
	usbArr[2] = Camera{"nikon"}
	fmt.Println(usbArr)

	for _, value := range usbArr{
		Computer{}.Working(value)
	}

	//computer := Computer{}
	//phone := Phone{}
	//camera := Camera{}
	//
	//computer.Working(phone)
	//computer.Working(camera)
}
