package main

import (
	"errors"
	"fmt"
)

func handleErr()  {
	err := recover()
	if err != nil{
		fmt.Println("err: ", err)
	}
}

func test(){
	defer handleErr()
	var num1, num2 = 1, 0
	res := num1 / num2
	fmt.Println(res)
}

func readConf(fName string)(err error){
	if fName == "config"{
		return nil
	}else {
		return errors.New("read config fail")
	}
}

func test2(){
	err := readConf("configs")
	if err != nil{
		panic(err)
	}
	fmt.Println("test2 done")
}

func main() {
	//test()
	//fmt.Println("test done")
	test2()
}