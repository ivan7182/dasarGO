package main

import (
	"fmt"
)

func main() {
	// name := "ivansv"

	// switch name {
	// case "ivan":
	// 	fmt.Println("hai ivan")
	// case "budi":
	// 	fmt.Println("hai budi")
	// default:
	// 	fmt.Println("kenalan dong")
	// }

	// switch panjang := len(name); panjang < 6 {
	// case true:
	// 	fmt.Println("nama terlalu pendek")
	// case false:
	// 	fmt.Println("nama terlalu panjang")
	// }

	// panjang := len(name)

	// switch {
	// case panjang > 6:
	// 	fmt.Println("nama terlalu panjang")
	// case panjang == 6:
	// 	fmt.Println("pas")
	// case panjang < 6:
	// 	fmt.Println("nama terlalu pendek")
	// }

	age := 11

	switch {
	case age < 12:
		fmt.Println("anak - anak")
	case age >= 12 && age <= 17:
		fmt.Println("remaja")
	case age <= 18 && age <= 64:
		fmt.Println("dewasa")
	case age <= 65:
		fmt.Println("lansia")

	}

}
