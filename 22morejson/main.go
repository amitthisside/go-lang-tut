package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Pricing  int      `json:"price"`
	Website  string   `json:"webportal"`
	Password string   `json:"-"` //- represents not showing the data in encoded JSON
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Let's learn more about JSON in Golang!")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	mySlice := []course{
		{"Python", 1299, "udemy.com", "python1212", []string{"python", "ai", "ml"}},
		{"DSA", 1099, "coursera.com", "dsa1812", []string{"data-structures", "algos"}},
		{"MERN", 1599, "codedamn.com", "mern2003", nil},
	}

	//encode the slice to JSON
	finalJson, err := json.MarshalIndent(mySlice, "", "\t ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
			{
				"coursename": "Python",
				"price": 1299,
				"webportal": "udemy.com",
				"tags": ["python","ai","ml"]
			}
	`)

	// var holdVals course

	// checkValid := json.Valid(jsonData)

	// if checkValid {
	// 	fmt.Println("JSON data is valid!")
	// 	json.Unmarshal(jsonData, &holdVals)
	// 	fmt.Printf("%#v\n", holdVals)
	// } else {
	// 	fmt.Println("JSON data is invalid!")
	// }

	//store JSON data in map

	var valueMap map[string]interface{}
	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON data is valid!")
		json.Unmarshal(jsonData, &valueMap)
		// fmt.Printf("%#v\n", valueMap)

		for k, v := range valueMap {
			fmt.Printf("Key : %v Value : %v Type : %T\n", k, v, v)
		}
	} else {
		fmt.Println("JSON data is invalid!")
	}

}
