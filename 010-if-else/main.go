package main

import "fmt"

func main() {
	fmt.Println("if-else")

	loginCount := 9

	if loginCount%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	//web request handling

	if num := 3; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Greater than or equal to 10")
	}
}
