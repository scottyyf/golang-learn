package main //表明这个go属于main包

import (
	"fmt"
	"os"
	"strings"
) //导入其他的包

func main() { //函数定义func，类似于def和shell中的function
	fmt.Println("hello, world")
	s, sep := "", ""
	for _, arg := range os.Args[1:]{//range返回的是一对结果
		s += sep + arg
		sep = " "
	}
	sep = strings.Join(os.Args[:], " ")
	for index, value := range os.Args {
		fmt.Println(index, value)
	}

	fmt.Println(sep)
}
