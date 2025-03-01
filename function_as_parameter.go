package main

import (
	"fmt"
	"strings"
)

// func sayHello(name string, filter func(string) string) {
// 	filteredName := filter(name)
// 	fmt.Println("hello", filteredName)
// }

// func spamFilter(name string) string {
// 	if name == "anjing" {
// 		return "***"
// 	} else {
// 		return name
// 	}
// }

// func main() {
// 	sayHello("anjing", spamFilter)
// }

// func prosesAngka(angka int, fungsiOperasi func(int) int) {
// 	hasil := fungsiOperasi(angka)
// 	fmt.Println(hasil)

// }

// func kaliDua(n int) int {
// 	return n * 2
// }

// func main() {
// 	prosesAngka(5, kaliDua)
// }

func prosesTeks(teks string, fungsiModifikasi func(string) string) {
	hasil := fungsiModifikasi(teks)
	fmt.Println(hasil)

}

func keKapital(s string) string {
	return strings.ToUpper(s)
}

func main() {
	prosesTeks("golang", keKapital)
}
