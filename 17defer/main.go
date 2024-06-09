package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in Golang!")
	//executes statement with defer at last
	//LIFO for multiple defers
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

//defer queues qill be created, multiple queues are possible

func myDefer() {
	defer fmt.Println()
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
