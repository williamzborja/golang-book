package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(non_comma("123"))
	fmt.Println(non_comma("12345"))
	fmt.Println(non_comma("123456"))
}

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	return comma(s[:n-3]) + "," + s[n-3:]
}

func non_comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteByte(s[i])

	}

	return buf.String()
}
