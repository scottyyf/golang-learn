package main

import "fmt"

func mapSlice(){
	//var a = make(map[string]string)
	var monsters = make([]map[string]string, 2)// 放两个妖怪

	var newMonster = make([]map[string]string, 1)
	newMonster[0] = map[string]string{
		"name": "masterMonkey",
		"age": "500",
	}

	if monsters[0] == nil{
		monsters[0] = make(map[string]string, 2)//两个属性
		monsters[0]["name"] = "masterCow"
		monsters[0]["age"] = "20"
	}

	if monsters[1] == nil{
		monsters[1] = make(map[string]string, 2)//两个属性
		monsters[1]["name"] = "masterRabbit"
		monsters[1]["age"] = "18"
	}

	//fmt.Println(monsters)
	currMaster := append(monsters, newMonster...)
	fmt.Println(currMaster)
}

func main() {
	mapSlice()
}
