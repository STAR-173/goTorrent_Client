package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackpal/bencode-go"
	// bencode "github.com/jackpal/bencode-go" // Available if you need it!
)

// Ensures gofmt doesn't remove the "os" encoding/json import (feel free to remove this!)
var _ = json.Marshal

type Metafile struct {
	Announce string   `bencode:"announce"`
	Info     Metainfo `bencode:"info"`
}

type Metainfo struct {
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

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
		contents, err := os.ReadFile(inputFile)

		var result Metafile

		err = bencode.Unmarshal(bytes.NewReader([]byte(contents)), &result)
		if err != nil {

			fmt.Println(err)

			return

		}

		h := sha1.New()

		bencode.Marshal(h, result.Info)

		fmt.Printf("Tracker URL: %v\n", result.Announce)
		fmt.Printf("Length: %v\n", result.Info.Length)
		fmt.Printf("Info Hash: %x\n", h.Sum(nil))

	default:
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
