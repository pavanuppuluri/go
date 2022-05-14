package main

import "fmt"

func main() {
   var marks int = 90

   switch marks {
      case 90: fmt.Printf("Excellent!\n" )  
      case 80: fmt.Printf("Well done\n" )  
      case 50,60,70 : fmt.Printf("You passed\n" ) 
      default: fmt.Printf("Failed\n" ); 
   }
    
}