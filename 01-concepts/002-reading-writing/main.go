package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Actual implementation
type Reader interface {
	Read(p []byte) (n int, err error)
}

// Common implementations

// *os.File: used to read from a file
// strings.Reader: creates a Reader from the provided string
// bufio.Reader: implements Reader for improved performance, using buffers for reading
// http.Request.Body: allows you to supply the request body via reader.
// bytes.Buffer: used to read from in-memory byte Buffers

func main() {

	f, err := os.Open("./cmd/001-Read-Write/letters.txt")
	if err != nil {
		panic(err)
	}
	count, err := countAlphabets(f)
	if err != nil {
		log.Panic("Error counting alphabets from file is", err)
	}
	fmt.Printf("The count of alphabets from file is %v\n", count)

	s := "This is another example"
	b := strings.NewReader(s)
	count, err = countAlphabets(b)
	if err != nil {
		log.Panic("Error counting alphabets from string\n", err)
	}
	fmt.Printf("The count of alphabets from string is %v", count)
}

func countAlphabets(r io.Reader) (int, error) {
	count := 0
	buffer := make([]byte, 1024)
	for {
		n, err := r.Read(buffer)
		for _, l := range buffer[:n] {
			if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
				count++
			}
		}
		if err == io.EOF {
			return count, nil
		}
		if err != nil {
			return 0, err
		}
	}
}
