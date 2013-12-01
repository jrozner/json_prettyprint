package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var unmarshaled interface{}

	if len(os.Args) != 2 {
		printUsage()
	}

	rawData, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(rawData, &unmarshaled)
	if err != nil {
		log.Fatal(err)
	}

	marshaled, err := json.MarshalIndent(unmarshaled, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(marshaled))
}

func printUsage() {
	fmt.Printf("Usage: %s <file>\n", os.Args[0])
	os.Exit(1)
}
