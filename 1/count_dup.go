package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { //读入数据则为true，否则为false，用来读取行数据
		content := input.Text()
		if content == "exit"{
			break
		}
		counts[content]++ //Text()获得具体值
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
		}
	}
	var s = []color.Color{color.White, color.Black}
	fmt.Println(s)

}
