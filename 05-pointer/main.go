package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println(ptr)

	mynumber := 23
	var myPtr *int = &mynumber
	fmt.Println(myPtr)
	fmt.Println(*myPtr)

	*myPtr = *myPtr * 2
	fmt.Println(mynumber)
}
