package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
  Name string
  Description string
}

func main() {
productJson := `{"name": "Laptop","description": "Dell Inspiron"}`
var product Product	
json.Unmarshal([]byte(productJson), &product)
fmt.Printf("Product Name: %s, Description: %s", product.Name, product.Description)
}