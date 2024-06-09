package main

import "fmt"

//outside any method, walrus (:=) operator for assigning variables is not allowed
//for eg. tokenId := 2319412
//the above initialization is invalid as it is not in a method and := operator is used

//constants can be decalred using const keyword
//making the first letter capital of the constant (or var.) makes it "public" in nature
const LoginAuthToken string = "fghuia9621hhka1" //this is public

func main() {
	var username string = "amitthisside"
	fmt.Println(username)
	fmt.Printf("Type of variable is : %T \n", username)

	var isAvailable bool = true
	fmt.Println(isAvailable)
	fmt.Printf("Type of variable is : %T \n", isAvailable)

	var points int = 420
	fmt.Println(points)
	fmt.Printf("Type of variable is : %T \n", points)

	var decimalValue float64 = 420.2346741213
	fmt.Println(decimalValue)
	fmt.Printf("Type of variable is : %T \n", decimalValue)

	//default values and aliases
	var anotherInt int
	fmt.Println(anotherInt) //default value : 0
	fmt.Printf("Type of variable is : %T \n", anotherInt)

	var anotherStr string
	fmt.Println(anotherStr) //default value : ""
	fmt.Printf("Type of variable is : %T \n", anotherStr)

	var isDefault bool
	fmt.Println(isDefault) //default value : false
	fmt.Printf("Type of variable is : %T \n", isDefault)

	var anotherFloat float64
	fmt.Println(anotherFloat) //default value : 0
	fmt.Printf("Type of variable is : %T \n", anotherFloat)

	//implicit type declaration
	//lexer automatically assigns type based upon value of variable
	var word = "rainingdogsandcats"
	fmt.Println(word)
	fmt.Printf("Type of variable is : %T \n", word)
	//word = 3 is not allowed as type is already set to string

	//no var declaration also works
	ranomVar := 312
	fmt.Println(ranomVar)

	//constant print
	fmt.Println(LoginAuthToken)
}
