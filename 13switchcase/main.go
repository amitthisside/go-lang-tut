package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to switch case program in Golang!")

	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Value is 1, you can open or movee 1 place")
	case 2:
		fmt.Println("Move 2 places")
	case 3:
		fmt.Println("Move 3 places")
		//fallthrough executes next case too
		fallthrough
	case 4:
		fmt.Println("Move 4 places")
		fallthrough
	case 5:
		fmt.Println("Move 5 places")
	case 6:
		fmt.Println("Move 6 plcaes and roll the dice again")
	default:
		fmt.Println("What's that??")
	}
}
