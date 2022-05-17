package main

import "fmt"

func main() {
	const (
		Sun = iota+1
		Mon = iota+1
		Tue = iota+1
		Wed = iota+1
		Thu = iota+1
		Fri = iota+1
		Sat = iota+1
		
	)

	fmt.Println(Sun, Mon, Tue, Wed, Thu, Fri, Sat)
}
