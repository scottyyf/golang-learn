package main

import "fmt"

type myint int
type funcType func(myint, myint) myint // 自定义数据类型

func getSum(num1 myint, num2 myint) myint {
	fmt.Printf("%T\n", num1)
	return num1 + num2
}

func funcTogether(getSum funcType, num1 myint, num2 myint) myint{
	return getSum(num1, num2)
}

func sumAndSub(n1 int, n2 int) (sum int, sub int){
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main(){
	//type myint int
	//type funcType func(myint, myint) myint
	/*
	var num1, num2 myint
	fmt.Scan(&num1, &num2)
	ret := funcTogether(getSum, num1, num2)
	fmt.Println(ret)

	 */
	var n1, n2 = sumAndSub(2, 1)
	fmt.Println(n1, " ",n2)
	fmt.Printf("%T", sumAndSub)
}