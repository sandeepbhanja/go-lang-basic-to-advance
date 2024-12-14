package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Rate our Pizza")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("You entered: %s", input)
	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil || newRating > 5 {
		fmt.Println("Invalid input. Please enter a number.")
	} else {
		fmt.Printf("Thank you for rating our pizza at %.2f out of 5!\n", newRating+1)
	}

}
