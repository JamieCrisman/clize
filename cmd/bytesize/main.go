package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("needs 1 argument: bytesize [number]\n")
		os.Exit(1)
	}
	bs := os.Args[1]
	if len(bs) == 0 {
		fmt.Printf("needs 1 argument: bytesize [number]\n")
		os.Exit(1)
	}
	b, err := strconv.Atoi(bs)
	if err != nil {
		fmt.Printf("'%v' was not integerable: %v\n", bs, err)
		os.Exit(1)
	}
	kb := float64(b) / 1_000
	mb := kb / 1_000
	gb := mb / 1_000

	fmt.Printf(strings.Repeat("%30s: %v\n", 4),
		"B", b,
		"KB", kb,
		"MB", mb,
		"GB", gb,
	)
}
