package main

import (
	"fmt"
)

func printJinZiTa(row int) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= row-i; j++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if i == row {
				fmt.Printf("%c", '*')
			} else if j == 1 || j == 2*i-1 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func printLingXing(row int) {
	for i := 1; i <= row; i++ {
		for j := 1; j <= row - i; j++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			if  j == 1 || j == 2*i-1{
				fmt.Printf("%c", '*')
			}else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	for i := row-1; i >= 1; i-- {
		for j := 1; j <= row - i; j++ {
			fmt.Printf(" ")
		}

		for j := 1; j <= 2*i-1 ; j++ {
			if j==1 || j== 2*i-1 {
				fmt.Printf("%c", '*')
			}else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func printNineNine(num int){
	for i := 1; i <= num; i++ {
		for j := 1; j <= i ; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func main() {
	printJinZiTa(4)
	printLingXing(5)
	printNineNine(5)
}
