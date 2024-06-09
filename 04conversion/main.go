package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Print("Please rate your recent pizza order : ")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating : ", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 32)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Adding 1 to your rating : ", numRating+1)
}
