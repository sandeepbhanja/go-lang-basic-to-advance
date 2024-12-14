package main

import "fmt"

func main() {
	fmt.Println("Maps")
	languages := make(map[string]int, 4)
	languages["Go"] = 2017
	languages["Python"] = 1991
	languages["Java"] = 1995
	languages["C++"] = 1985
	fmt.Println(languages)
	delete(languages, "Go")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("%s: %d",key,value)
	}
}
