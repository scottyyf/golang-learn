package main

import "fmt"

func breakLevel() {
label0:
	for i := 0; i < 4; i++ {
	label1:
		for j := 0; j < 10; j++ {
			if j == 5 {
				continue label0
			}
			if j == 8 {
				break label1
			}
			fmt.Printf("%d\t", j)
		}

	}
}

func sumTo20() {
	var ret int
	for i := 0; i < 101; i++ {
		ret += i
		fmt.Printf("%d\t", i)
		if ret == 55 {
			break
		}
	}
	fmt.Println()
}

func loginCheck() {
	var name, password string
	for i := 0; i < 3; i++ {
		fmt.Printf("inset name: ")
		fmt.Scanln(&name)
		fmt.Printf("inset password: ")
		fmt.Scanln(&password)
		if name != "zwj" || password != "888" {
			fmt.Printf("failed, %d time to retry\n", 2-i)
			continue
		}
		fmt.Println("success")
		break
		}
	}


func main() {
	breakLevel()
	//sumTo20()
	//loginCheck()
}
