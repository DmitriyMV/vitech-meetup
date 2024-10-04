package main

func main() {
	start := 0
	end := 10

	rangeNum := func(yield func(int) bool) {
		for i := start; i < end; i++ {
			yield(i)
		}
	}

	for i := range rangeNum {
		println(i)

		break
	}
}
