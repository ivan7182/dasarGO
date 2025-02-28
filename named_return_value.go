package main

import "fmt"

func getCompleteName() (firstname, lastName, address string) {
	firstname = "ivan"
	lastName = "satria"
	address = "tegal"

	return firstname, lastName, address
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)

}
