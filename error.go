package main

import "errors"
import "fmt"

func divide(value1 int, value2 int)(int, error) {
   if(value2 == 0){
      return 0, errors.New("Divide by zero error")
   }
   return value1/value2, nil
}
func main() {
   result, err:= divide(9,0)

   if err != nil {
      fmt.Println(err)
   } else {
      fmt.Println(result)
   }
   
}