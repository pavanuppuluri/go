package main

import "fmt"

func main() {

	rates1 := [...]float64{
		1:   1.5, // index: 1
		2:   2.4, // index: 2
		0:   0.5, // index: 0
	}
	
	fmt.Println(rates1)
	
	rates2 := [...]float64{
		// index 1 to 4 empty

		5:   1.5, // index: 5
		2.5, //      index: 6
		0:   0.5, // index: 0
	}

	fmt.Println(rates2)

	// above array literal equals to this:
	//
	// rates := [7]float64{
	// 	0.5,
	// 	0.,
	// 	0.,
	// 	0.,
	// 	0.,
	// 	1.5,
	//  2.5,
	// }
}