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

	// age := 11

	// switch {
	// case age < 12:
	// 	fmt.Println("anak - anak")
	// case age >= 12 && age <= 17:
	// 	fmt.Println("remaja")
	// case age <= 18 && age <= 64:
	// 	fmt.Println("dewasa")
	// case age <= 65:
	// 	fmt.Println("lansia")

	// }

	// moon := 6

	// switch moon {
	// case 1:
	// 	fmt.Println("januari")
	// case 2:
	// 	fmt.Println("februari")
	// case 3:
	// 	fmt.Println("maret")
	// case 4:
	// 	fmt.Println("april")
	// case 12:
	// 	fmt.Println("desember")
	// default:
	// 	fmt.Println("error")
	// }

	var angka int

	fmt.Print("masukan angka:")
	_, err := fmt.Scan(&angka)

	if err != nil {
		fmt.Println("error")
		return
	}

	switch {
	case angka%2 == 0:
		fmt.Printf("%d adalah bilangan genap.\n", angka)
	default:
		fmt.Printf("%d adalah bilangan ganjil.\n", angka)
	}

}
