package main

import (
	"fmt"
	"golang-learn/2022_last/project3/utils"
)

func main() {
	var b utils.B
	b.X.Name = "tom"
	b.X.Age = 18
	b.Name = "scott"

	b.X.SayOk()
	b.X.Hello()

	b.C.SayOk()
	b.X.Wataxi()
	b.X.Hello()
	b.Setint(10.0)
	fmt.Println(b)
}
