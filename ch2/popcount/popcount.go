package main

import (
	"fmt"
)

// pc[i] is the population count of i.

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	sum := 0

	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

func main() {
	fmt.Println(PopCount(255))
}
