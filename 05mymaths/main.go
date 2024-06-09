package main

import (
	"fmt"

	"math/rand"
	// "crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in GoLang!")

	//random number using math/rand lib
	// rand.Seed(12) -> dperecated
	fmt.Println(rand.Intn(5))

	//random number using crypto/rand lib
	// fmt.Println(rand.Reader, big.NewInt(5))
}
