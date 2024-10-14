package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackpal/bencode-go"
	// bencode "github.com/jackpal/bencode-go" // Available if you need it!
)

// Ensures gofmt doesn't remove the "os" encoding/json import (feel free to remove this!)
var _ = json.Marshal

// Example:
// - 5:hello -> hello
// - 10:hello12345 -> hello12345

func main() {

	command := os.Args[1]

	switch command {
	case "decode":
		// Uncomment this block to pass the first stage

		bencodedValue := os.Args[2]

		decoded, err := bencode.Decode(bytes.NewReader([]byte(bencodedValue)))
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))

	case "info":
		inputFile := os.Args[2]
		contents, _ := os.ReadFile(inputFile)
		decoded, _ := bencode.Decode(bytes.NewReader([]byte(contents)))

		fmt.Println("Tracker URL: ", decoded.(map[string]any)["announce"])
		fmt.Println("Length: ", decoded.(map[string]any)["info"].(map[string]any)["length"])

	default:
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
