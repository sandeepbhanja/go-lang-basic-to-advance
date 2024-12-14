package main

import "fmt"

const LoginToken string = "lucifer"

func main() {
	var username string = "Sandeep"
	fmt.Println("username:" + username)
	fmt.Printf("Variables is of type: %T\n", username)

	var isLoggedIn bool = false
	fmt.Println("isLoggedIn:" + fmt.Sprintf("%t", isLoggedIn))

	var smallVal int8 = 127
	fmt.Printf("smallVal is of type: %d\n", smallVal)

	var smallFloat float32 = 132.23467890
	fmt.Printf("smallFloat is of type: %f\n", smallFloat)

	//default values and aliases
	var anotherVariable int
	fmt.Printf("anotherVariable is of type: %d\n", anotherVariable)

	//implicit values
	var website = "xyz.com"
	fmt.Println("website:" + website)

	//no var style
	// numberOfUsers := 3000
	fmt.Println("numberOfUsers:" + fmt.Sprintf("%s", LoginToken))
}
