package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case")
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(3) + 1
	switch randomNumber {
	case 1:
		fmt.Printf("You got a %d\n", randomNumber)
	case 2:
		fmt.Printf("You got a %d\n", randomNumber)
	case 3:
		fmt.Printf("You got a %d\n", randomNumber)
	default:
		fmt.Println("Default case")
	}
}
