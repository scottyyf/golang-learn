package main

import "fmt"

func gotoPrint(ok bool) string{
	var str = "中华人民共和国中央政府's"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\t", str2[i])
	}
	for index, val := range str{
		fmt.Printf("%d, %c\n", index, val)
	}

	if ok{
		return "OK"
	}else {
		return "NO"
	}

}

func main() {
	var ret = gotoPrint(false)
	fmt.Println(ret)
}
