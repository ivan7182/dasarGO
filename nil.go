package main

import "fmt"

func Describe(i interface{}) {
	if i == nil {
		fmt.Println("interface is nil")
	} else {
		fmt.Println("interface has a value")
	}
}

func CheckNumberPointer(p *int) {
	if p == nil {
		fmt.Println("pointer is nil")
	} else {
		fmt.Println("Pointer value", *p)
	}
}

func main() {
	Describe(nil)
	Describe("Golang")
	Describe(200)

	var p *int

	CheckNumberPointer(p)
	x := 42
	p = &x
	CheckNumberPointer(p)
}
