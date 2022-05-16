package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	chOdd := make(chan int)
	chEven := make(chan int)

	go odd(chOdd)
	go even(chEven)

	for _, value := range intSlice {
		if value%2 != 0 {
			chOdd <- value
		} else {
			chEven <- value
		}
	}

}

func odd(ch <-chan int) {
	for v := range ch {
		fmt.Println("ODD :", v)
	}
}

func even(ch <-chan int) {
	for v := range ch {
		fmt.Println("EVEN:", v)
	}
}
