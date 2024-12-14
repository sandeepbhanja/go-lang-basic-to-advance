package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var url string = "https://example.com"

func main() {
	fmt.Println("WebRequest")
	response, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Printf("Data received: %s\n", string(databytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
