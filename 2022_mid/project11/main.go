package main

import (
	"fmt"
	"math"
)

func stdScore(class int, students int) {
	//cl, st := int(class), int(students)
	var scores [3][5]int
	for i := 0; i < class; i++ {
		for stu := 0; stu < students; stu++ {
			fmt.Scanln(&scores[i][stu])
		}
	}
	fmt.Println(scores)
}

func scoreGiver() {
	var teacher [8]float64
	for i := 0; i < len(teacher); i++ {
		fmt.Scanln(&teacher[i])
	}

	fmt.Println(teacher)
	for j := 0; j < len(teacher) - 1; j++ {
		for i := 0; i < len(teacher)-j-1; i++ {
			if teacher[i] > teacher[i+1] {
				tmp := teacher[i]
				teacher[i] = teacher[i+1]
				teacher[i+1] = tmp
			}
		}
	}
	fmt.Println(teacher)
	var total float64 = 0
	for i := 1; i < len(teacher) - 1; i++ {
		total += teacher[i]
	}
	avg := total / float64(len(teacher)-2)
	fmt.Println(avg)
	if (avg - teacher[0]) > teacher[len(teacher)-1] - teacher[0]{
		fmt.Println("bad score", teacher[0])
	} else if  (avg - teacher[0]) == teacher[len(teacher)-1] - teacher[0]{
		fmt.Println("bad score", teacher[0], teacher[len(teacher)-1])
	}else {
		fmt.Println("bad score", teacher[len(teacher)-1])
	}

	bestScore := teacher[0]
	for i := 1; i < len(teacher); i++ {
		if math.Abs(teacher[i] - avg) < math.Abs(bestScore - avg){
			bestScore = teacher[i]
		}
	}
	fmt.Println(bestScore)
}

func main() {
	//stdScore(3, 5)
	scoreGiver()
}