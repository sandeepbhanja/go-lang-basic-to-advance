package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Sandeep"
	fmt.Printf("Hello , %s", welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	//comma ok || err err
	input, _ := reader.ReadString('\n')
	fmt.Printf("Name: %s", input)

}
