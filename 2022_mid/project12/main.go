package main

import "fmt"

func testMap() {
	var student = make(map[string]map[string]string)
	student["tom"] = map[string]string{"id": "01", "sex": "female"}
	student["jerry"] = map[string]string{"id": "02", "sex": "female"}
	student["snoby"] = map[string]string{"id": "03", "sex": "male"}
	for name, m := range student {
		fmt.Printf("%s\t", name)
		for key, value := range m {
			fmt.Printf("%s: %s\t",key, value)
		}
		fmt.Println()
	}
}

func main() {
	testMap()
}
