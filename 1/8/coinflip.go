package main

var test = 1

func Signum(x int) int {
	switch  {
	case x >0:
		return 1
	default:
		return 0
	case x<0:
		return -1

	}
}
/*

 */

func main()  {
	ret := Signum(5)
	println(ret)
	println(&test)
	println(*&test)
}
