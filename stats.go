package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

const (
	defaultSource = "./data/fridge_temperature.json"
)

func main() {
	file := flag.String("file", "", "Source file data")
	flag.Parse()

	// Get source file from command or
	// set default
	var source string
	if *file == "" {
		source = defaultSource
	} else {
		source = *file
	}

	// Fridge component
	fridge := FridgeComponent{
		Source: source,
	}
	output, err := fridge.Temperature()
	if err != nil {
		log.Println(err)
	}

	data, err := json.Marshal(output)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
}
