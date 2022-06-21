package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {

	const url = "localhost:8000"
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	log.Println("listen in port :8000")
	log.Fatalln(http.ListenAndServe(url, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	log.Println(count)

	mu.Unlock()

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Count:", count)
}
