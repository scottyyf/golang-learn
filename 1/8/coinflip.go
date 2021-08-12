package main

func Signum(x int) int {
	switch  {
	case x >0:
		return 1
	default:
		return 'res'
	case x<0:
		return -1

	}
}

func main()  {
	ret := Signum(5)
	println(ret)
}
