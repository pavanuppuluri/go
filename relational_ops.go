package main

import "fmt"

func main() {
   var a int = 10
   var b int = 20

   if( a == b ) {
      fmt.Printf("a is equal to b\n" )
   } else {
      fmt.Printf("a is not equal to b\n" )
   }
   if ( a < b ) {
      fmt.Printf("a is less than b\n" )
   } else {
      fmt.Printf("a is not less than b\n" )
   } 
   if ( a > b ) {
      fmt.Printf("a is greater than b\n" )
   } else {
      fmt.Printf("a is not greater than b\n" )
   }
   
   if ( a <= b ) {
      fmt.Printf("a is either less than or equal to  b\n" )
   }
   if ( b >= a ) {
      fmt.Printf("b is either greater than  or equal to b\n" )
   }
}