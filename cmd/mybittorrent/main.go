package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	// bencode "github.com/jackpal/bencode-go" // Available if you need it!
)

// Ensures gofmt doesn't remove the "os" encoding/json import (feel free to remove this!)
var _ = json.Marshal

// Example:
// - 5:hello -> hello
// - 10:hello12345 -> hello12345

func decodeString(b string, st int) (x interface{}, i int, err error) {
	var l int

	i = st

	for i < len(b) && b[i] >= '0' && b[i] <= '9' {
		l = l*10 + (int(b[i]) - '0')
		i++
	}
	if i == len(b) || b[i] != ':' {
		return "", st, fmt.Errorf("Bad String")
	}

	i++

	if i+1 > len(b) {
		return "", st, fmt.Errorf("Bad String")
	}

	x = b[i : i+l]
	i += l

	return x, i, nil
}

func decodeInt(b string, st int) (x int, i int, err error) {
	i = st
	i++

	if i == len(b) {
		return 0, st, fmt.Errorf("Bad Int")
	}

	neg := false

	if b[i] == '-' {

		neg = true

		i++

	}

	for i < len(b) && b[i] >= '0' && b[i] <= '9' {

		x = x*10 + (int(b[i]) - '0')

		i++

	}

	if i == len(b) || b[i] != 'e' {
		return 0, st, fmt.Errorf("Bad Int")
	}

	i++

	if neg {
		x = -x
	}

	return x, i, nil
}

func decodeList(b string, st int) (l []interface{}, i int, err error) {
	i = st
	i++

	l = make([]interface{}, 0)

	for {
		if i >= len(b) {
			return nil, st, fmt.Errorf("Bad List")
		}

		if b[i] == 'e' {
			break
		}

		var x interface{}

		x, i, err = decode(b, i)
		if err != nil {
			return nil, i, err
		}

		l = append(l, x)
	}

	return l, i, nil
}

func decode(b string, st int) (x interface{}, i int, err error) {
	if st == len(b) {
		return nil, st, io.ErrUnexpectedEOF
	}

	i = st

	switch {
	case b[0] == 'i':
		return decodeInt(b, i)
	case b[i] == 'l':
		return decodeList(b, i)
	case b[i] >= '0' && b[i] <= '9':
		return decodeString(b, i)
	default:
		return nil, st, fmt.Errorf("unexpected value %q", b[i])
	}
}

func main() {

	command := os.Args[1]

	if command == "decode" {
		// Uncomment this block to pass the first stage

		bencodedValue := os.Args[2]

		decoded, err := decodeBencode(bencodedValue)
		if err != nil {
			fmt.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		fmt.Println("Unknown command: " + command)
		os.Exit(1)
	}
}
