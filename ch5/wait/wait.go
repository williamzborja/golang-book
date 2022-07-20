package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://gopl.io"

	log.SetPrefix("wait: ")
	log.SetFlags(0)

	if err := WaitForServer(url); err != nil {
		log.Fatalln("Site is down: ", err)
	}
}

func WaitForServer(url string) error {
	const timeout = time.Minute * 1
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}

		log.Printf("server not responding (%s); retrying...", err)
		fmt.Printf("Sleeping %d", time.Second<<uint(tries))
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
