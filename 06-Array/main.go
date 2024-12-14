package main

import "fmt"

func main() {
	fmt.Println("Array")

	var fruitList [3]string
	fruitList[0] = "Apple"
	fruitList[2] = "Banana"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
}
