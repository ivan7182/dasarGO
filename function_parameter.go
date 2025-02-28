package main

import "fmt"

func Hello(firtName string, lastName string) {
	fmt.Println("hello", firtName, lastName)
}
func main() {
	Hello("ivan", "satria")
}
