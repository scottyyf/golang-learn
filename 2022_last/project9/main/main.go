package main

import (
	"golang-learn/2022_last/project9/utils"
)


func twoSum(nums []int, target int) []int {
	var data = map[int]int{}
	//res := make([]int, 2)
	var res = new([]int)
	for index, value := range nums{
		another := target - value
		ret, ok := data[another]
		if ok{
			*res = []int{index, ret}
			return *res
		}else{
			data[value] = index
		}
	}
	return *res
}

func main() {
	var familyAccount = utils.FamilyAccount{
		ExitTag: false,
		Balance: 10000.0,
		IsFirst: true,
		Details: "\ttotal\tinput\tnote",
	}
	//fmt.Println(familyAccount)
	familyAccount.MainMenu()
	//nums := []int{2,7,11,15}
	//target := 9
	//x := twoSum(nums, target)
	//fmt.Println(x)
}
