package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(names string) string{
	return func(name string) string {
		if strings.HasSuffix(name, suffix){
			//fmt.Println("no suffix")
			return name
		}else {
			//fmt.Println(name)
			name += suffix
			return name
		}
	}
}

func main() {
	var suffix = ".jpg"
	var name = "十大元帅jpg"
	endSuffix := makeSuffix(suffix)
	//endSuffix(name)
	ret := endSuffix(name)
	fmt.Println(ret)
	fmt.Println(name)
}