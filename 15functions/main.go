package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang!")
	fmt.Println(adder(3, 12))
	fmt.Println(multiAdder(1, 5, 12, 5, 23))
	res, str := multiValueReturner(3, 1)
	fmt.Println(res, str)
}

func adder(a int, b int) int {
	return a + b
}

func multiAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func multiValueReturner(a int, b int) (int, string) {
	return (a + b), "return string value"
}
