package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in Golang!")
	//No inheritance; No super or parent concept
	amit := Student{21, "Amit Mishra", true}
	fmt.Println(amit)
	fmt.Printf("Detailed view : %+v\n", amit)
	fmt.Printf("Name is %v\n", amit.Name)
	amit.IsPassed()
	amit.ChangeName()
	fmt.Println("Name :", amit.Name)
}

type Student struct {
	RollNo int
	Name   string
	Passed bool
}

//functions under structs are known as methods
func (s Student) IsPassed() {
	fmt.Println("Passing status is", s.Passed)
}

func (s Student) ChangeName() {
	s.Name = "Jon Doe"
	fmt.Println("Name :", s.Name)
}
