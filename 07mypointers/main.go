package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers class!")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr) //value is nil

	n := 21
	var ptr = &n
	fmt.Println("Adress of pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println(n)
}
