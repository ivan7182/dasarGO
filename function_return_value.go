package main

import "fmt"

func getHello(name string) string {
	hello := "hello " + name
	return hello
}

func main() {
	result := getHello("ivan")
	fmt.Println(result)
}
