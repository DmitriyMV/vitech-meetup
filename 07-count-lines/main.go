package main

import (
	"bufio"
	"io"
	"iter"
	"os"
)

//go:generate /bin/sh -c "echo 'Hello\nworld!\nThis is ViTech' | go run $GOFILE"

func main() {
	n := 0

	for _, err := range readLine(bufio.NewReader(os.Stdin)) {
		if err != nil {
			println(err.Error())

			return
		}

		n++
	}

	println(n)
}

func readLine(rdr io.Reader) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		scanner := bufio.NewScanner(rdr)

		for scanner.Scan() {
			if !yield(scanner.Text(), nil) {
				break
			}
		}

		if err := scanner.Err(); err != nil {
			yield("", err)
		}
	}
}
