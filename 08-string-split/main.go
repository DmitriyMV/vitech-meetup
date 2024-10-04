package main

import "strings"

func main() {
	const str = "Hello, World!, I'm on ViTech meetup"

	for _, part := range strings.Split(str, ",") {
		println(strings.TrimLeft(part, " "))
	}
}
