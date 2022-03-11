package main

import (
	"fmt"
)

func homeWork1(c byte) {
	switch c {
	case 'a', 'b', 'c', 'd', 'e':
		fmt.Printf("%c", c - 32)
		fmt.Println()
	default:
		fmt.Println("other")
	}
}

func homeWork2(score float32) {
	switch {
	case score > 100:
		fmt.Println("invalid score: ", score)
	case score >= 60:
		fmt.Println("passed")
	default:
		fmt.Println("injected")
	}
}

func homeWork3(month uint8) {
	/*
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Run()
	 */
	switch month / 3 {
	case 1:
		fmt.Println("spring")
	case 2:
		fmt.Println("summer")
	case 3:
		fmt.Println("autumn")
	case 4, 0:
		if month > 12{
			fmt.Println("invalid")
		}else {
			fmt.Println("winter")
		}

	default:
		fmt.Println("invalid")
	}

}

func homeWork4() {
	fmt.Println()
}

func main() {
	homeWork1('c')
	homeWork2(60.01)
	homeWork3(13)
	homeWork4()
}
