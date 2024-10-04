package main

import (
	"iter"
	"strings"
)

func main() {
	for part := range split("Hello, World!, I'm on ViTech meetup", ",") {
		println(strings.TrimLeft(part, " "))
	}
}

func split(s, sep string) iter.Seq[string] {
	return func(yield func(string) bool) {
		if sep == "" {
			panic("split: empty separator")
		}

		n := strings.Count(s, sep) + 1

		if n > len(s)+1 {
			n = len(s) + 1
		}
		n--
		i := 0
		for i < n {
			m := strings.Index(s, sep)
			if m < 0 {
				break
			}

			if !yield(s[:m]) {
				return
			}

			s = s[m+len(sep):]
			i++
		}

		if !yield(s) {
			return
		}
	}
}
