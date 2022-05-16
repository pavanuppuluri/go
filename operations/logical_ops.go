package main

import "fmt"

func main() {
   var a bool = true
   var b bool = false
   if ( a && b ) {
      fmt.Printf("Condition is true\n" )
   }
   if ( a || b ) {
      fmt.Printf("Condition is true\n" )
   }
   
   if ( !(a && b) ) {
      fmt.Printf("Condition is true\n" )
   }
}