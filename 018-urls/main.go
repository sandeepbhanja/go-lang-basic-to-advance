package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://jsonplaceholder.typicode.com/posts?userId=1"

func main() {
	fmt.Println(myUrl)

	//parsing
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	for key, value := range qparams {
		fmt.Printf("%s=%s", key, value)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "jsonplaceholder.typicode.com",
		Path:     "/posts",
		RawQuery: "userId=1",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
