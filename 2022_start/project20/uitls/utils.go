package utils

import "fmt"

var Age int = 100
var Name string = "tom"

// main.go use it

func init(){
	fmt.Println("utils init called")
	Age = 101
	Name = "Tom"
}