package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "hi go language is one of the best programming language"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]


	for _, q := range query {

		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// find the first word then break
				// the nested loop
				break
			}
		}

	}
}