package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Hero struct {
	Name string
	Age int
}

type HeroSlice []Hero

func (s HeroSlice) Len() int {
	return len(s)
}

func (s HeroSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s HeroSlice) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func main() {
	//h1 := Hero{"n1", 20}
	//h2 := Hero{"n2", 10}
	//h3 := Hero{"n3", 15}
	var hs HeroSlice
	for i:= 0; i< 100; i++{
		hero := Hero{
			Name: fmt.Sprintf("hero-%d", rand.Intn(100)),
			Age: rand.Intn(100),
		}
		hs = append(hs, hero)
	}
	for _, value := range hs{
		fmt.Println(value)
	}

	//hs := HeroSlice{h1, h2, h3}
	//fmt.Println(hs)
	time.Sleep(2000000000)
	fmt.Println("====================================")
	sort.Sort(hs)
	for _, value := range hs{
		fmt.Println(value)
	}
}
