package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name   string `json:"Name"`
	RollNo int    `json:"Roll"`
}

func main() {

	p := Person{
		Name:   "Sathwik",
		RollNo: 66,
	}

	f, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error Creating file: ", err)
		panic(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(p)

	if err != nil {
		fmt.Println("Error encoding person", err)
		panic(err)
	}
}
