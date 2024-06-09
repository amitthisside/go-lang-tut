package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://courses.chaicode.com/learn/account/signup?signInToPay=193290&priceId=135451"

func main() {
	fmt.Println("URLs in Golang!")
	fmt.Println(myurl)

	//parsing our url
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Path)

	params := result.Query()
	fmt.Println(params)

	for key, val := range params {
		fmt.Printf("Key : %v Value : %v\n", key, val)
	}

	//construct a URL from parts
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "amit.dev",
		Path:    "/projects/go-projects",
		RawPath: "user=temp",
	}

	constructedUrl := partsOfUrl.String()
	fmt.Println(constructedUrl)

}
