package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to file handling in Golang!")

	//let's create and write to a txt file
	content := "This string goes right into the file (I really hope it does :/)"

	file, err := os.Create("./mytextfile.txt")
	checkNilErr(err)

	lengthOfFile, err := file.WriteString(content)
	checkNilErr(err)

	fmt.Println("Length of file :", lengthOfFile)
	defer file.Close() //defer close ensures that the file is closed at the end of execution of program

	readFromFile("./mytextfile.txt")
}

func readFromFile(filename string) {
	content, err := os.ReadFile(filename) //content is received in byte format, always do conversions accordingly
	checkNilErr(err)
	fmt.Println("File content :", string(content))

}

// common nil error checker function
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
