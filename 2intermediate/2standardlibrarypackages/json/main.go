package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Ninja struct {
	Name  string `json:"ninja_name"`
	Level int    `json:"ninja_level"`
}

func main() {
	// Package json implements functionalities involving the
	// encoding and decoding between JSON and Go values
	tommy := Ninja{Name: "Tommy", Level: 9001}
	fmt.Println(tommy)

	// Marshal returns the JSON encoding of v
	jsonByteSlice, err := json.Marshal(tommy)
	fmt.Println(string(jsonByteSlice), err)

	// Unmarshal converts JSON back Go
	var tommyJr Ninja
	err = json.Unmarshal(jsonByteSlice, &tommyJr)
	fmt.Println(tommyJr)
	var tommySr Ninja
	jsonString := `{"ninja_name":"Tommy Sr", "ninja_level":9001}`
	err = json.Unmarshal([]byte(jsonString), &tommySr)
	fmt.Println(tommySr, err)
	fmt.Println()

	// Go object <-> io.Reader/io.Writer using json.Decoder/json.Encoder
	var tommyAgain Ninja
	err = json.NewDecoder(os.Stdin).Decode(&tommyAgain)
	fmt.Println(tommyAgain, err)
	err = json.NewEncoder(os.Stdout).Encode(tommyAgain)
}
