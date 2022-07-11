package main

import (
	"fmt"
	"time"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var a [3]int
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d, %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	// literal arrays

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}

	fmt.Println(q, r)

	q2 := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q2)

	t := time.Hour * 2

	fmt.Printf("%[1]v\n", t)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	fmt.Println(symbol, RMB, symbol[RMB])

	r2 := [...]int{99: -1}
	fmt.Println(r2)

	months := [...]string{1: "January", "Febrary", "March", "April", "May", "June", "July",
		"August", "September", "October", "November",
		12: "December"}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	fmt.Println(cap(summer), len(summer))
	fmt.Println(summer[:7])
	fmt.Println(cap(summer[:7]), len(summer[:7]))
	text := "hello"
	fmt.Printf("%T\n", text[1:])

	var stack []int

	stack = append(stack, 3)   // push
	top := stack[len(stack)-1] // top
	fmt.Println(top)

	stack = stack[:len(stack)-1] // pop

	fmt.Println(stack)
}
