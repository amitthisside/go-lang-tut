package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang!")

	mySlice := []string{"June", "Jan", "May", "April", "Dec"}
	fmt.Println(mySlice)
	// for i := 0; i < len(mySlice); i++ { //++i is invalid
	// 	fmt.Println(mySlice[i])
	// }

	// for idx := range mySlice { //this is returning index values to the idx var
	// 	fmt.Println(mySlice[idx])
	// }

	// for idx, val := range mySlice { //now both idx and vals are returnes
	// 	fmt.Printf("Index : %v Value : %v\n", idx, val)
	// }

	randomVal := 1
	for randomVal < 11 {
		//can also use break, continue and goto
		// if randomVal == 6 {
		// 	break
		// }
		// if randomVal == 5 {
		// 	randomVal++ //updation when using continue is important, otherwise infinite loop
		// 	continue
		// }

		if randomVal == 7 {
			goto myLabel
		}
		fmt.Println(randomVal)
		randomVal++
	}

myLabel:
	fmt.Println("this is a random goto label")
}
