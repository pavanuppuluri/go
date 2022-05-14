package main

import "fmt"

func main() {
   var a int = 21
   var c int

   c -=  a
   fmt.Printf("-=  Operator Example, Value of c = %d\n", c )

   c +=  a
   fmt.Printf("+= Operator Example, Value of c = %d\n", c )

   c *=  a
   fmt.Printf("*= Operator Example, Value of c = %d\n", c )

   c /=  a
   fmt.Printf("/= Operator Example, Value of c = %d\n", c )

   c  = 200; 

   c <<=  2
   fmt.Printf("<<= Operator Example, Value of c = %d\n", c )

   c >>=  2
   fmt.Printf(">>= Operator Example, Value of c = %d\n", c )

   c &=  2
   fmt.Printf("&= Operator Example, Value of c = %d\n", c )

   c ^=  2
   fmt.Printf("^= Operator Example, Value of c = %d\n", c )

   c |=  2
   fmt.Printf("|= Operator Example, Value of c = %d\n", c )
}