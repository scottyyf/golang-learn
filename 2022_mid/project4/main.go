package main

import "fmt"

func getDayFromMonth()(day int){
	return 30
}

func chikenInfo(){
	/*
	n1, n2, n3, n4, n5, n6 := 1.0, 2.0, 3.0 ,4.0, 5.0, 6.0
	avg := (n1 + n2+n3+n4+n5+n6)/6.0
	sum := n1+n2+n3+n4+n5+n6
	fmt.Println(avg, " -- ", sum)
	 */
	var chicken [6]float64
	//chicken = {1.0, 2.0, 3.0 ,4.0, 5.0, 6.0}
	for i := 0; i < len(chicken); i++ {
		
	}

	fmt.Println(len(chicken))
}

func main() {
	days := getDayFromMonth()
	fmt.Println(days)
	chikenInfo()
}
