package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("!Files")
	content := "Sandeep Bhanja"
	file, err := os.Create("./example.txt")
	io.WriteString(file, content)
	checkNilErr(err)
	file1, err := os.Open("example.txt")
	checkNilErr(err)

	readFile(file1.Name())

	defer file.Close()
	defer file1.Close()
}

func readFile(fileName string) {
	databytes, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println(string(databytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
