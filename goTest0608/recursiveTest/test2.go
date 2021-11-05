package main

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n - 1) + Fibonacci(n - 2)
}

func main()  {
	for i := 1; i < 10; i++ {
		println(Fibonacci(i))
	}
}
