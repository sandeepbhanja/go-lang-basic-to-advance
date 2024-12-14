package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")
	fruitList := []string{"Apple"}
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 100
	highScores[1] = 95
	highScores[2] = 90
	highScores[3] = 85
	highScores = append(highScores, 999)
	fmt.Println(len(highScores))

	sort.Ints(highScores)

	fmt.Println(highScores)

	courses := []string{"C++", "C", "Python", "JS"}
	fmt.Println(courses)
	index := 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
