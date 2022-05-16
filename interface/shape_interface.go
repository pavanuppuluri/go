package main

import (
   "fmt" 
   "math" 
)

type Shape interface {
    area() float64
}

/* define a circle */
type Circle struct {
   radius float64
}

/* define a rectangle */
type Rectangle struct {
   width,height float64
}

/* define a method for circle */
func(circle Circle) area() float64 {
   return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle */
func(rectangle Rectangle) area() float64 {
   return rectangle.width * rectangle.height
}

func calculateArea(shape Shape) float64 {
    return shape.area()
}

func main(){
   circle := Circle{radius:5}
   rectangle := Rectangle{width:10, height:20}
   fmt.Printf("Circle area: %f", calculateArea(circle))
   fmt.Printf("Circle area: %f", calculateArea(rectangle))
}
