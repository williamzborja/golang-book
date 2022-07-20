package main

import (
	"fmt"
	"strings"
)

func add1(r rune) rune { return r + 1 }

func double(x int) (result int) {
	defer func() { fmt.Printf("%d + %d = %d\n", x, x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

// func Parser(input string) (s *Syntax, err error) {
// 	defer func() {
// 		if p := recover(); p != nil {
// 			err = fmt.Errorf("internal error :%v", p)
// 		}
// 	}()
// 	//
// }

func main() {

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"

	fmt.Println(triple(3))
}
