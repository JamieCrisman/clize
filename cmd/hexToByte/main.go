package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		txt = strings.TrimPrefix(txt, "0x")
		bytes, err := hex.DecodeString(txt)
		if err != nil {
			fmt.Printf("could not decode string: %v", err)
			os.Exit(1)
		}

		fmt.Printf("%30s: `%v`\n%30s: %#v\n%30s: %+v\n",
			"stringed", string(bytes),
			"go style", bytes,
			"bytes", bytes,
		)
	}
}
