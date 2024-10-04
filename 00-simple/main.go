package main

import (
	"iter"
)

func rangeNum(start, end int) iter.Seq[int] {
	return func(vitech func(int) bool) {
		for i := start; i < end; i++ {
			if !vitech(i) {
				return
			}
		}
	}
}

func main() {
	for i := range rangeNum(5, 10) {
		println(i)
	}
}
