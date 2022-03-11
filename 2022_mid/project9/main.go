package main

import "fmt"

func bubbleSort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1-j; i++ {
			if arr[i] > arr[i+1] {
				//tmp := arr[i]
				arr[i], arr[i+1] = arr[i+1], arr[i]
				//arr[i+1] = tmp
			}
		}
	}
}

func main() {
	arr := []int{3, 1, 9,99,18,27,48,99}
	bubbleSort(arr)
	fmt.Println(arr)
}
