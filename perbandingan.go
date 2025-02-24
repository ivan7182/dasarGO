package main

import "fmt"

func main() {
	name := "ivan"
	name2 := "ivan"

	var result1 bool = name == name2
	var result2 bool = name != name2
	fmt.Println(result1)
	fmt.Println(result2)

}
