package main

import (
	"fmt"
)

func marryHim(height int32,
	money float32,
	handsome bool) bool {
	if height > 180 && money > 50 && handsome {
		fmt.Println("say yes")
		return true
	} else if height > 180 || money > 50 || handsome {
		fmt.Println("say yes so so")
		return true
	} else {
		fmt.Println("fuck off")
		return false
	}
}

func switchBranch(data int) {
	//switch {
	//case data > 80:
	//	fmt.Println("large than 80")
	//case data > 70:
	//	fmt.Println("large than 70")
	//default:
	//	fmt.Println("less than 70")
	//
	//}
	switch {
	case true:
		fmt.Println("80")
		fallthrough
	case data > 70:
		fmt.Println("70")
	case data > 60:
		fmt.Println("60")
	default:
		fmt.Println("default")
	}
}

func typeSwitch(){
	var x interface{}
	var y int = 32
	x = y
	switch x.(type) {
	case nil:
		fmt.Println("nil object")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("i do not know")
	}

}

func main() {
	var height int32 = 90
	var money float32 = 10
	var handsome bool = true
	var ret bool
	//fmt.Scan(&height, &money, &handsome)
	ret = marryHim(height, money, handsome)
	fmt.Println(ret)
	switchBranch(0)
	typeSwitch()
}
