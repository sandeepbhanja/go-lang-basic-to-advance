package main

import "fmt"

func main() {
	fmt.Println("loops")

	days := make([]string, 8)
	days[0] = "Sunday"
	days[1] = "Monday"
	days[2] = "Tuesday"
	days[3] = "Wednesday"
	days[4] = "Thursday"
	days[5] = "Friday"
	days[6] = "Saturday"
	days[7] = "Sunday"

	// for k := 0; k < 7; k++ {
	// 	fmt.Printf("%d: %s\n", k+1, days[k])
	// }

	for index, day := range days {
		fmt.Printf("%d: %s\n", index, day)
	}
}
