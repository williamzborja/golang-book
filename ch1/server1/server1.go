package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	log.Println("listening in port :8000")
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalln(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path: %q", r.URL.Path)
}
