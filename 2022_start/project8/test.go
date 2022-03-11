package main

import "fmt"

func twoToTen() {
	// 1011 = 1* 2^0 + 1*2^1 + 0*2^2 + 1*2^3
	var t = "11"
	var ret int
	for i := len(t) - 1; i >= 0; i-- {
	}
	fmt.Println(ret)
}

func eightToTen() {
	// 0123 = 3*8^0 + 2*8^1 + 1 *8^2 + 0*8^3

}

func sixteenToTen() {
	// 0x1234 = 4 * 16^0 + 3*16^1 + 2*16^2 + 1*16^3
}

func tenToTwo() {
	// 1234
	// ret = 1234 / 2  个位
	// ret = ret / 2  第二位
}

func tenToEight() {
	// 01234
	// ret = 1234 / 8
	// ret = ret /8
}

func tenToSixteen() {
	//	0x123F
	//	ret = 123F --> 12315 / 16
	//	ret = ret /16 ...
}

func twoToEight() {
	//	每三位一组,转成对应8进制
	//	11 010 101 --> 3 2 5
}

func twoToSixteen() {
	//	每四位一组,转成对应16进制
	//	1101 0101 --> 13 5 --> 0xd5
}

func eightToTwo() {
	//	将8进制的每一位转成3位对应的二进制
	//	0237 --> 010 011 111
}

func sixteenToTwo() {
	// 将16进制的每一位转成4位对应的二进制
	// 0x237 --> 0010 0011 0111
}

func main() {
	//
	twoToTen()
}
