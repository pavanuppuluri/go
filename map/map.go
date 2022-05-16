package main

import "fmt"

func main() {
   var personEmailMap map[string]string
   /* create a map*/
   personEmailMap = make(map[string]string)
   
   /* insert key-value pairs in the map*/
   personEmailMap["Pavan"] = "pavan@gmail.com"
   personEmailMap["Kavya"] = "kavya@gmail.com"
   personEmailMap["Lasya"] = "lasya@gmail.com"
   personEmailMap["Pandu"] = "pandu@gmail.com"
   
   /* print map using keys*/
   for name := range personEmailMap {
      fmt.Println("Email of",name," is ",personEmailMap[name])
   }
   
   /* test if entry is present in the map or not*/
   capital, ok := personEmailMap["Lasya"]
   
   /* if ok is true, entry is present otherwise entry is absent*/
   if(ok){
      fmt.Println("Lasya email", capital)  
   } else {
      fmt.Println("Lasya email not present") 
   }
   
   /* delete an entry */
   delete(personEmailMap,"Pavan");
   
      /* print map using keys*/
   for name := range personEmailMap {
      fmt.Println("Email of",name," is ",personEmailMap[name])
   }
   
}
