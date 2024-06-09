package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps tutorial!")

	shortForms := make(map[string]string)

	shortForms["JS"] = "Javascript"
	shortForms["PY"] = "Python"
	shortForms["RB"] = "Ruby"

	fmt.Println(shortForms)
	fmt.Println("JS is short form for", shortForms["JS"])

	delete(shortForms, "RB")
	fmt.Println(shortForms)

	//iterating through the map
	for key, value := range shortForms {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
