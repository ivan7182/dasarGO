package main

import "fmt"

func main() {
	// person := map[string]string{
	// 	"name":   "ivan",
	// 	"addres": "tegal",
	// }

	// fmt.Println(person["name"])
	// fmt.Println(person["addres"])
	// fmt.Println(person)

	// book := map[string]string{
	// 	"title":  "golang",
	// 	"Author": "ivan",
	// 	"ups":    "salah",
	// }
	// fmt.Println(book)

	// delete(book, "Author")
	// fmt.Println(book)

	// // Menambahkan dan Mengakses Elemen dalam Map
	// person1 := map[string]int{
	// 	"alice": 30,
	// 	"bob":   25,
	// }
	// person1["iput"] = 90
	// fmt.Println(person1)

	// //Mengubah Nilai dalam Map
	// score := map[string]int{
	// 	"math":    90,
	// 	"sciense": 85,
	// 	"english": 92,
	// }
	// fmt.Println("sebelum diubah", score)

	// score["sciense"] = 88
	// fmt.Println(score)

	// //Menghapus Elemen dalam Map
	// inventory := map[string]int{
	// 	"apple":  5,
	// 	"banana": 3,
	// 	"cherry": 10,
	// }

	// delete(inventory, "banana")
	// fmt.Println(inventory)

	sembako := map[string]string{
		"name":  "gula",
		"jenis": "sembako",
	}

	fmt.Println(sembako)
	delete(sembako, "name")
	fmt.Println(sembako)

}
