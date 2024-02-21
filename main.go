package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func Factorial(n int) (result int) {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	return n * Factorial(n-1)
}

func main() {
	fmt.Println(Add(2, 3))
}
