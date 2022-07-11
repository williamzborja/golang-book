package main

import (
	"fmt"
	"sort"
)

func main() {
	// age := make(map[string]int)

	ages := map[string]int{
		"alice":   31,
		"charlie": 32,
	}
	fmt.Println(ages)
	fmt.Println(ages["alice"])

	delete(ages, "charlie")
	fmt.Println(ages)

	ages["william"] = 29
	ages["angui"] = 25

	names := make([]string, 0, len(ages))

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	fmt.Println("Sorted map")
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
