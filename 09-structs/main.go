package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// no inheritance in golang ; No super or parent
	// no classes
	sandeep := Person{"Sandeep", 22, "xyz@a.com", false}
	fmt.Println(sandeep)
	fmt.Printf("%+v", sandeep)
}

type Person struct {
	Name   string
	Age    int
	Email  string
	Status bool
}
