package main

import "fmt"

func main() {
	name := "ivans"

	// if name == "budi" {
	// 	fmt.Println("hello")
	// } else if name == "harjo" {
	// 	fmt.Println("hallo harjo")
	// } else {
	// 	fmt.Println("kenalan dong")
	// }

	if panjang := len(name); panjang > 6 {
		fmt.Println("nama terlalu panjang")
	} else if panjang == 5 {
		fmt.Println("pas")
	} else {
		fmt.Println("nama terlalu pendek")
	}

	// nilai := 6

	// if nilai > 7 {
	// 	fmt.Println("good")
	// } else if nilai == 6 {
	// 	fmt.Println("tingkatkan Lagi")
	// } else {
	// 	fmt.Println("remidi")
	// }
}
