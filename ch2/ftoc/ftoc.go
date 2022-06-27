package main

import "fmt"

func main() {
	const freezingF, boolingF = 32.0, 212.0

	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boolingF, fToC(boolingF))

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
