package main

import "fmt"

func main() {
	var name = "ivan"
	var e = name[3]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
