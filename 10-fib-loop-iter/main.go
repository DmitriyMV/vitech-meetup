package main

import "iter"

func fibIter(n int, yield func(int) bool) {
	f1 := 1
	f2 := 0

	for i := 0; i < n; i++ {
		if !yield(f2) {
			return
		}

		f1, f2 = f2, f1+f2
	}

	yield(f2)
}

func fib(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		fibIter(n, yield)
	}
}

func main() {
	for n := range fib(20) {
		println(n)
	}
}
