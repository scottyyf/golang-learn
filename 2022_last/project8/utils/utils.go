package utils

import "fmt"

type Student struct {
}

func TypeJudge(items ...interface{}) {
	for i, item := range items {
		switch item.(type) {
		case string:
			fmt.Printf("%d is a string, %v\n", i, item)
		case bool:
			fmt.Printf("%d is a bool, %v\n", i, item)
		case float32:
			fmt.Printf("%d is a float32, %v\n", i, item)
		case Student:
			fmt.Printf("%d is a student, %v\n", i, item)
		case *Student:
			fmt.Printf("%d is *student, %v\n", i, item)
		default:
			fmt.Printf("%d is %T, value %v\n", i, item, item)
		}
	}
}
