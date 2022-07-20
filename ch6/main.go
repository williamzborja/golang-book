package main

import (
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{mapping: make(map[string]string)}

var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

func lookup(key string) string {
	mu.Lock()
	defer mu.Unlock()
	v := mapping[key]
	return v
}

func main() {

}
