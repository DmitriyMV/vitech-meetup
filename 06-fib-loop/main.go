package main

func fib(n int) int {
	f1 := 1
	f2 := 0

	for i := 0; i < n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}

func main() {
	println(fib(10))
}
