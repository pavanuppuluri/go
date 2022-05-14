package main

import "fmt"

func counter() func() int {

    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    increment := counter()

    fmt.Println(increment())
    fmt.Println(increment())
    fmt.Println(increment())
    fmt.Println(increment())

}