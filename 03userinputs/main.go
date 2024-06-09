package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMsg := "Welcome to the start of the user input program!"
	fmt.Println(welcomeMsg)

	//input from bufio, input type Stdin from os
	//create a reader or input var
	reader := bufio.NewReader(os.Stdin)

	//this reader can be used to take inputs
	fmt.Print("Enter a rating for the urban clap service : ")

	//comma ok || error ok syntax
	//comma seperates two vals, first one holds the val if everything goes right, second one holds value if anyhting goes wrong
	//go doesn't have try-catch, so its like try, catch vals
	//for not using any value amongst the two, use underscore (_) instead

	ratingInput, _ := reader.ReadString('\n') //ratingInput stores val if everything is alright, otherwise the default val _ is used for error part
	fmt.Println("Thanks for your rating : ", ratingInput)
}
