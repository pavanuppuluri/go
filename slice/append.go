package main

import "fmt"

func main() {
	var todo []string

	todo = append(todo, "go")
	todo = append(todo, "react")
	todo = append(todo, "java", "mongodb")

	// you can also append a slice to another slice using ellipsis: ...
	
	tomorrow := []string{"kafka", "aws"}
	todo = append(todo, tomorrow...)
	

	fmt.Println(todo);
}