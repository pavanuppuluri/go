package main

import (
    "fmt"
)

func main() {

    numArray := []int{9,11,90,6,8,21}
        
    
    for _, value := range numArray {
        if(value%2 == 0) {
         fmt.Printf("First even number in a given array is %d\n", value)   
         break;
        }
    }
    
}