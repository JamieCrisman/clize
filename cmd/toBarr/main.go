package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Printf("%30s: `%v`\n%30s: %#v\n%30s: %+v\n",
			"stringed", string(txt),
			"byte arr", []byte(txt),
			"input", txt,
		)
	}
}
