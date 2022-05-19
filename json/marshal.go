package main

import (
	"encoding/json"
	"fmt"
)



type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions string	    `json:"permissions"`
}

func main() {
	users := []user{
		{"pavan", "1234", "read-only"},
		{"sai", "42", "admin"},
		{"sruthilasya", "66", "write"},
	}

	
	out, err := json.MarshalIndent(users, "", "\t")
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(string(out))
}
