package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {

		if !strings.HasPrefix(arg, "http://") && !strings.HasPrefix(arg, "https://") {
			arg = "http://" + arg
		}
		resp, err := http.Get(arg)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(resp.Status)

		//		_, err = io.Copy(os.Stdout, resp.Body)

		/*		if err != nil {
					fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				}
		*/
	}
}
