package it

import (
	"bufio"
	"io"
	"iter"
)

// ReadLines produces a sequence of lines from the input.
func ReadLines(r io.Reader) iter.Seq[string] {
	return func(yield func(string) bool) {
		if closer, ok := r.(io.Closer); ok {
			defer closer.Close()
		}
		s := bufio.NewScanner(r)
		s.Split(bufio.ScanLines)
		for s.Scan() {
			if !yield(s.Text()) {
				return
			}
		}
	}
}
