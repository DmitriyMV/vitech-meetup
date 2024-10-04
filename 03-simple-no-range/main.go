package main

func main() {
	start := 0
	end := 10

	rangeNum := func(yield func(int) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}

	rangeNum(func(i int) bool {
		println(i)

		return true
	})
}
