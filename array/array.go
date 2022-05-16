package main

import (
    "fmt"
)

func main() {

    numArray := []int{1, 2, 3,4,5}
        

    /* Approach 1 */
    fmt.Println(numArray)
    
    /* Approach 2 */
    for _, value := range numArray {
    fmt.Printf("%d\n", value)
    }
    
}
