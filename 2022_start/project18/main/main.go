package main

import "fmt"

func test(num int) {
	if num < 1 {
		return
	}

	fmt.Printf("%d\t", num)
	num--

	test(num)
}

func feb(num int) int {
	if num == 1 || num == 2 {
		return 1
	}
	left, right, ret := 1, 1, 0
	for i := 3; i <= num; i++ {
		ret = left + right
		left = right
		right = ret
		fmt.Printf("%d\t", ret)
	}
	fmt.Println()
	return ret
}

func febRachio(num int) int {
	if num < 3 {
		return 1
	}
	ret := febRachio(num-1) + febRachio(num-2)
	return ret
}

func getSum(n1 int, n2 int) {
	ret := n1 + n2
	fmt.Println(ret)
}

func originNum(n int) int {
	if n > 10 || n < 1 {
		return 0
	} else if n == 10 {
		return 1
	}
	ret := (originNum(n+1) + 1) * 2
	return ret
}

func main() {
	//test(10)
	//getSum(10, 20)
	//x := febRachio(10)
	//fmt.Println(x)
	x := originNum(2)
	fmt.Println(x)
}
