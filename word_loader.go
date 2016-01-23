package gopher_typer

import (
	"bufio"
	"io"
)

func NewWordLoader(r io.Reader) []string {
	var words []string
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words
}
