package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateLeft(s []int, rotations int) {
	reverse(s[:rotations])
	reverse(s[rotations:])
	reverse(s)
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	rotateLeft(s, 2)
	fmt.Println(s)
}
