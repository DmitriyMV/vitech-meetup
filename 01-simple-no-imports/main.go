package main

func rangeNum(start, end int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
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
