package main

import "fmt"

func orderFind(name int, orders []int) (index int) {
	//# first sort this slice
	for outer := 0; outer < len(orders)-1; outer++ {
		for inner := 0; inner < len(orders)-1-outer; inner++ {
			if orders[inner] > orders[inner+1] {
				tmp := orders[inner]
				orders[inner] = orders[inner+1]
				orders[inner+1] = tmp
			}
		}
	}
	fmt.Println(orders)
	// binary find
	left, right := 0, len(orders) - 1
	for {
		if left > right{
			break
		}
		mid := (left + right) / 2
		if orders[mid] == name{
			return mid
		}else if orders[mid] < name{
			left = mid + 1
		}else{
			right = mid - 1
		}
	}
	return -1
}

func main() {
	var orders = []int{1, 5, 3, 8, 0, 2, 30, 19, 22, 55}
	var heroName int
	fmt.Scanln(&heroName)
	index := orderFind(heroName, orders)
	fmt.Println(index)
}
