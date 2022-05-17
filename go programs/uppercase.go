package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper(os.Args[1]))
}