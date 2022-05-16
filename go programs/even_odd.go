package main

import "fmt"

func main() {
   /* local variable definition */
   var a int = 100;
 
   /* check the boolean condition */
   if( a%2 == 0 ) {
      fmt.Printf("a is even \n" );
   } else {      
      fmt.Printf("a is odd \n" );
   }
}