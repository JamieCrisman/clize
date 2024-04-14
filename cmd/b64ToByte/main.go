package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		bytes, err := base64.StdEncoding.DecodeString(txt)
		if err != nil {
			fmt.Printf("could not decode: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%30s: `%v`\n%30s: %#v\n%30s: %+v\n",
			"stringed", string(bytes),
			"go style", bytes,
			"bytes", bytes,
		)
	}
}
