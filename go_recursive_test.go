package main

import (
	"fmt"
	"testing"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func TestResursive(t *testing.T) {
	fmt.Println(fact(6))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return n * fib(n-1)
	}

	fmt.Println(fib(5))
}
