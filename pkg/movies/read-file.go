package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// type Stuff struct {
// 	Name string
// 	Jorb string
// 	Age string
// }

func ReadJsonFile() {
	jsonFile, err := os.Open("./data.json")

	if err != nil {
		fmt.Println("error reading json file", err)
	}

	// Defer closing the file, so we can parse it later.
	defer jsonFile.Close()

	// Read the opened json file as a byte array
	byteValue, _ := io.ReadAll(jsonFile)

	var stuffs map[string]interface{}

	json.Unmarshal(byteValue, &stuffs)

	fmt.Println(stuffs)
}
