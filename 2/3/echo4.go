package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit taiiling newline") //n是指针m,是bool的，默认false
var sep = flag.String("s", ", ", "separator") // s是指针,string的，默认", "



func main()  {
	flag.Parse() //初始化
	fmt.Print(strings.Join(flag.Args(), *sep)) //连接string
	if !*n{
		fmt.Println()
	}
}
