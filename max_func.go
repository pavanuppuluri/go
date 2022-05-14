package main

import "fmt"

func main() {
 
   var a int = 111
   var b int = 222

   fmt.Printf( "Max value is : %d\n", max(a, b) )
}

/* function returning the max of given two numbers */
func max(num1, num2 int) int {
   

   if (num1 > num2) {
      return num1
   } else {
      return num2
   }
}