package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Golang!")
	//No inheritance; No super or parent concept
	amit := Student{21, "Amit Mishra", true}
	fmt.Println(amit)
	fmt.Printf("Detailed view : %+v\n", amit)
	fmt.Printf("Name is %v and passing status is %v", amit.Name, amit.Passed)
}

type Student struct {
	RollNo int
	Name   string
	Passed bool
}
