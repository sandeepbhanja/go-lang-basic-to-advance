package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 Monday")) //MM-DD-YYY

	createdDate := time.Date(2020, time.December, 25, 20, 35, 30, 0, time.UTC)
	fmt.Println(createdDate)
}
