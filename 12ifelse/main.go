package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to if-else control flow in Golang!")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	age, _ := strconv.Atoi(strings.TrimSpace(input))
	var msg string

	if age >= 18 {
		msg = "Valid voter"
	} else {
		msg = "Invalid voter"
	} // this syntax is sensitive to spacing and new lines

	fmt.Println(msg)

	//another syntax can initialize a value and then use it
	if n := 5; n%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
