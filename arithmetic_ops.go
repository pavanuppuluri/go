package main

import "fmt"

func main() {

   var a int = 10
   var b int = 20
   var c int

   c = a + b
   fmt.Printf("Result of + is %d\n", c )
   
   c = a - b
   fmt.Printf("Result of - is %d\n", c )
   
   c = a * b
   fmt.Printf("Result of * is %d\n", c )
   
   c = a / b
   fmt.Printf("Result of / is %d\n", c )
   
   c = a % b
   fmt.Printf("Result is %d\n", c )
   
   a++
   fmt.Printf("Result of ++ is %d\n", a )
   
   a--
   fmt.Printf("Result of -- is %d\n", a )
}