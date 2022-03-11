package main

import (
	"fmt"
	"golang-learn/2022_last/project8/utils"
)

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

func (p Phone) Start() {
	fmt.Println("phone ", p.Name, "start")
}

func (p Phone) Stop() {
	fmt.Println("phone ", p.Name, "stop")
}

func (p Phone) Call() {
	fmt.Println("phone ", p.Name, "called")
}

func (c Camera) Start() {
	fmt.Println("camera start ", c.Name)
}

func (c Camera) Stop() {
	fmt.Println("camera ", c.Name, "stop")
}

type Computer struct{

}

func (c Computer) Working(usb Usb)  {
	usb.Start()
	if u, ok := usb.(Phone);ok{
		u.Call()
	}
	usb.Stop()
}

func main() {
	//var usbArr [3]Usb
	//usbArr[0] = Phone{"huawei"}
	//usbArr[1] = Phone{"xiaomi"}
	//usbArr[2] = Camera{"zhenhua"}
	//fmt.Println(usbArr)
	//var computer Computer
	//for _, value := range usbArr{
	//	computer.Working(value)
	//}
	value := byte(10)
	v2 := 123
	v3 := utils.Student{}
	v4 := new(utils.Student)
	utils.TypeJudge(value, v2, v3, v4)
}
