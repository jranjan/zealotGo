package main

import "fmt"

func main() {
	fmt.Printf("Factorial of %d is %d", 6, fact(6))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

