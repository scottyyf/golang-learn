package main

import (
	"fmt"
	"math/rand"
	"time"

	//"strconv"
)

func test() {
	var upperCase [26]byte
	var startCase byte = 'A'
	//start := 'A' + 1
	//fmt.Printf("%T", startCase)
	//fmt.Println()
	for i := 0; i < len(upperCase); i++ {
		//upperCase[i] = startCase + i
		upperCase[i] = startCase + byte(i)
		//fmt.Printf("%c", upperCase[i])
	}
	//fmt.Println(upperCase)
	//x := 'B' + 1
	//strconv.FormatInt(x, 10)
	//fmt.Printf("%x", x)
	maxIndex := 0
	for i := 0; i < len(upperCase); i++ {
		if upperCase[i] > upperCase[maxIndex] {
			maxIndex = i
		}
	}

	var total int
	for _, value := range upperCase {
		total += int(value)
	}
	var avg = float64(total) / float64(len(upperCase))

	fmt.Println(total, avg)
	//fmt.Println()
	//fmt.Printf("%d --: %c", maxIndex,upperCase[maxIndex])
}

func test1() {
	var ret [5]int
	for i := 0; i < len(ret); i++ {
		rand.Seed(time.Now().UnixNano())
		ret[i] = rand.Int()
	}
	fmt.Println(ret)
	left, right := 0, len(ret) - 1
	for{
		if left >= right{
			break
		}

		tmp := ret[left]
		ret[left] = ret[right]
		ret[right] = tmp

		left ++
		right --
	}
	fmt.Println(ret)
}

func main() {
	//test()
	test1()
}
