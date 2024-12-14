package main

import "fmt"

func main() {
	fmt.Println("methods")
	sandeep := User{"sandeep", 22, true}
	sandeep.GetStatus()
}

type User struct {
	Name   string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Printf("%+v\n", u)
}
