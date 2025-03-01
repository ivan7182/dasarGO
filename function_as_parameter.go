package main

import (
	"fmt"
	"strings"
)

type filter func(string) string
type fungsiOperasi func(int) int

func sayHello(name string, filter filter) {
	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "***"
	} else {
		return name
	}
}

func prosesAngka(angka int, fungsiOperasi fungsiOperasi) {
	hasil := fungsiOperasi(angka)
	fmt.Println(hasil)

}

func kaliDua(n int) int {
	return n * 2
}

func prosesTeks(teks string, fungsiModifikasi func(string) string) {
	hasil := fungsiModifikasi(teks)
	fmt.Println(hasil)

}

func keKapital(s string) string {
	return strings.ToUpper(s)
}

func cekUmur(umur int, validator func(int) bool) {
	if validator(umur) {
		fmt.Println("dewasa")
	} else {
		fmt.Println("belum cukup dewasa")
	}
}

func minimal18(umur int) bool {
	return umur <= 18
}

func hitungDisckon(harga float64, dickonCalculator func(float64) float64) {
	hasil := dickonCalculator(harga)
	fmt.Println(hasil)
}

func diskon10(harga float64) float64 {
	return harga * 0.9
}

func prosesNama(nama string, formatter func(string) string) {
	hasil := formatter(nama)
	fmt.Println(hasil)
}

func tambahMR(nama string) string {
	return "MR" + nama
}

func main() {
	// sayHello("anjing", spamFilter)
	//prosesTeks("golang", keKapital)
	// cekUmur(10, minimal18)
	hitungDisckon(100000, diskon10)
	prosesNama("ivan", tambahMR)
	// prosesAngka(5, kaliDua)
}
