package main

import "fmt"

func main() {
	fmt.Println("Welcome to learning arrays!")

	//arrays are not much used in go lang, not much functionalities in arrays

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[2] = "Strawberry"
	fruits[3] = "Papaya"

	fmt.Println("Fruit list:", fruits)
	fmt.Println("Length of fruit list:", len(fruits)) //len is 4, even when only 3 items are inserted

	var veg = [5]string{"Potato", "Onions", "Beans"}
	fmt.Println("Vegetable List:", veg)
	fmt.Println("Length of veg list:", len(veg))
}
