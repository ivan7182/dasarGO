package main

import "fmt"

func welcome(name string) string {
	return "hello " + name
}

func main() {

	a := welcome
	fmt.Println(a("ivan"))

}
