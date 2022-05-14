package main

import "fmt"

func main() {
   var n [10]int 
   var i,j,sum int

        
   for i = 0; i < 10; i++ {
      n[i] = i
   }
   

   for j = 0; j < 10; j++ {
		sum = sum+n[j];
	}
   fmt.Printf("Sum = %d\n", sum )
}