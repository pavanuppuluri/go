package main

import (
	"fmt"
	"sort"
)

func main() {
	
	cities := []string{"Chennai", "Hyderabad", "Bangalore"}
	sort.Strings(cities)
	fmt.Println(cities)

}
