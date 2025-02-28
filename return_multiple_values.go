package main

import "fmt"

func getBiodata() (string, string, int) {
	return "ivan", "satria", 25
}

func main() {
	firstName, _, age := getBiodata()
	fmt.Println(firstName, age)
}
