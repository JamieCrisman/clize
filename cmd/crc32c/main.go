package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "", "the name of the file to read")
	flag.Parse()

	var input io.ReadCloser
	if len(filename) != 0 {
		filepath, err := filepath.Abs(filename)
		if err != nil {
			filepath = filename
		}
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Printf("could not open file: %v\n", err)
			os.Exit(1)
		}
		input = file
		defer input.Close()
	} else {
		input = os.Stdin
	}

	buffer := make([]byte, 4096)
	reader := bufio.NewReader(input)

	table := crc32.MakeTable(crc32.Castagnoli)
	hash := crc32.New(table)

	totalBytes := 0
	for {
		n, err := reader.Read(buffer)
		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Printf("error reading file: %v\n", err)
			input.Close()
			//nolint:gocritic
			os.Exit(1)
		}
		if n == 0 {
			break
		}
		totalBytes += n
		hash.Write(buffer[:n])
	}

	fmt.Printf("%20s: %v\n%20s: %v\n",
		"crc32 (castagnoli)", hash.Sum32(),
		"bytes", totalBytes,
	)
}
