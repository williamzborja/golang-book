package main

import (
	"fmt"
	"math"
	"os"
)

func t() error {
	if f, err := os.Open("main.go"); err != nil {
		return err
	} else {
		// f and err are visible here too
		stat, err := f.Stat()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(stat.Name(), stat.Size(), "bytes")
		f.Close()
	}
	return nil
}

var x uint8 = 1<<1 | 1<<5
var y uint8 = 1<<1 | 1<<2

func main() {

	fmt.Println(math.MaxFloat64)
	fmt.Printf("%08b :%v\n", x, x)
	fmt.Printf("%08b :%v\n", y, y)

	fmt.Printf("%08b :%v\n", x&y, x&y)
	fmt.Printf("%08b :%v\n", x|y, x|y)
	fmt.Printf("%08b :%v\n", x^y, x^y)
}
