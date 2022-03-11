package main

import (
	"encoding/json"
	"fmt"
)

type Cats struct {
	Name  string
	Age   int
	Score [5]float64
	ptr   *int
	slice []int
	map1  map[string]string
}

func test(name string, cat Cats) bool {
	return false
}

type Point struct {
	X string
	Y string
}

type point struct {
	X string `json:"name"`
	Y string `json:"age"`
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	var a point = point{"时间反对","robiit"}
	ret, _ := json.Marshal(a)
	fmt.Println(string(ret))
}
