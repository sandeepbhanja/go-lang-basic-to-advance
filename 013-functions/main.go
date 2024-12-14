package main

import "fmt"

func main() {
	fmt.Println("Functions")
	greet()

	sumValue := proAdder(1,2,3,4,5,6,)
	fmt.Println(sumValue)
}

func greet() {
	fmt.Println("Hello Lucifer")
}

func add(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
