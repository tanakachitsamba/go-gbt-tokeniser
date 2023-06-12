package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// define a struct to hold the JSON output
type EncoderOutput struct {
	Tokens []int `json:"tokens"`
	Count  int   `json:"count"`
}

func main() {
	// the string to be encoded
	//str := "This is an example sentence to try encoding out on!"

	const str = process.argv[2] // get the input string from the command line arguments

	// run the JavaScript file with Node.js
	cmd := exec.Command("node", "encoder.js", str)

	// get the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Command failed with error: %v. Output: %s", err, output)
	}

	// parse the JSON output
	var result EncoderOutput
	err = json.Unmarshal(output, &result)
	if err != nil {
		log.Fatalf("Failed to parse JSON output: %v", err)
	}

	// print the encoded string and token count
	fmt.Printf("Encoded tokens: %v\n", result.Tokens)
	fmt.Printf("Token count: %d\n", result.Count)
}
