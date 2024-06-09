package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Golang!")

	var fruits = []string{"Apple", "Guava", "Peach"}
	fmt.Printf("Type of fruit list: %T\n", fruits)

	fruits = append(fruits, "Cherry", "Raspberry")
	fmt.Println("Fruit list:", fruits)
	fmt.Println("Length of fruit list:", len(fruits))

	slicedList := fruits[1:4]
	fmt.Println(slicedList)

	highScores := make([]int, 4)
	highScores[0] = 314
	highScores[1] = 110
	highScores[2] = 451
	highScores[3] = 700
	// highScores[4] = 910 // show error
	highScores = append(highScores, 901, 417)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//remove a value from slice based on index
	var courses = []string{"react", "rust", "golang", "python", "java", "mongodb"}
	fmt.Println(courses)
	var indexToRemove int = 3
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)
	fmt.Println(courses)
}
