package gopherTyper

import (
	"bufio"
	"io"
)

func newWordLoader(r io.Reader) []string {
	var words []string
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words
}
