package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := `
	
	This string contains spaces.It should be printed with out spaces.   
	`

	fmt.Println(strings.TrimSpace(msg))
}