package main

import "fmt"

func fbn(n int)[]int{
	var fb = make([]int, n)
	fb[0], fb[1] = 1, 1
	for i := 2; i < len(fb); i++ {
		fb[i] = fb[i-1] + fb[i-2]
	}
	return fb
}

func main(){
	fb := fbn(10)
	fmt.Println(fb)
}
