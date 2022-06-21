package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	args := os.Args[1:]

	if len(args) == 0 {
		countLines(os.Stdin, counts)
		return
	}

	for _, arg := range args {
		f, err := os.Open(arg)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}

	for key, value := range counts {
		if value > 1 {
			fmt.Println(value, key)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[f.Name()+":"+input.Text()]++
	}
}
