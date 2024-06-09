package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://amitthisside.github.io/tindog/"

func main() {
	fmt.Println("Let's handle web requests in Golang!")

	response, err := http.Get(url)

	checkNilErr(err)

	fmt.Printf("Type of response : %T\n", response) //response returns original reference, not a copy of it
	defer response.Body.Close()                     //caller's responsibility to close the response

	body, err := io.ReadAll(response.Body)
	checkNilErr(err)

	content := string(body)
	fmt.Println("Body is :\n", content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
