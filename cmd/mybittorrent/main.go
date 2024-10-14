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
		var result Metafile

		inputFile := os.Args[2]
		contents, _ := os.ReadFile(inputFile)
		decoded, _ := bencode.Decode(bytes.NewReader([]byte(contents)))

		h := sha1.New()

		bencode.Marshal(h, result.Info)

		fmt.Printf("Tracker URL: %v\n", decoded.(map[string]any)["announce"])
		fmt.Printf("Length: %v\n", decoded.(map[string]any)["info"].(map[string]any)["length"])
		fmt.Printf("Info Hash: %x\n", h.Sum(nil))

	default:
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
