package main

import "fmt"

func main() {
	fmt.Println(gcd(25, 5))
	fmt.Println(fib(25))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}
