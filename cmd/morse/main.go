package main

import (
	"fmt"
	"github.com/vkuragin/morse"
	"os"
	"strings"
)

func main() {
	action := os.Args[1]
	args := strings.Join(os.Args[2:], " ")

	m, err := morse.New()
	if err != nil {
		panic(err)
	}

	switch action {
	case "encode":
		encoded := m.Encode(args)
		fmt.Printf("%s -> %s\n", args, encoded)
	case "decode":
		decoded := m.Decode(args)
		fmt.Printf("%s -> %s\n", args, decoded)
	default:
		panic(fmt.Sprintf("Unknown action: %s\n", action))
	}

}
