package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb!")

	// const mygeturl = "http://localhost:7000/get"
	// PerformGetRequest(mygeturl)

	// const myposturl = "http://localhost:7000/post"
	// PerformPostJsonRequest(myposturl)

	const mypostformurl = "http://localhost:7000/postform"
	PerformPostFormRequest(mypostformurl)
}

func PerformGetRequest(myurl string) {
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Content Length :", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	//alternative way using strings.Builder
	var responseString strings.Builder
	byteLen, _ := responseString.Write(content)
	fmt.Println("Byte Length :", byteLen)
	fmt.Println(responseString.String())

	// stringContent := string(content)

	// fmt.Println(stringContent)
}

func PerformPostJsonRequest(myurl string) {
	//generate fake json payload
	responseBody := strings.NewReader(`
		{
			"name":"Amit Mishra",
			"roll no":21,
			"status":"enrolled"
		}
	`)

	response, err := http.Post(myurl, "application/json", responseBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	var stringResponse strings.Builder
	byteLen, _ := stringResponse.Write(content)

	fmt.Println("Content/Byte Length :", byteLen)
	fmt.Println(stringResponse.String())
}

func PerformPostFormRequest(myurl string) {
	//form data
	data := url.Values{}
	data.Add("fname", "Amit")
	data.Add("lname", "Mishra")
	data.Add("email", "code@amit.in")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
