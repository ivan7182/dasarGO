package main

import "fmt"

func main() {
	nilaiAkhir := 74
	nilaiAbsen := 74

	var lulusNilaiAkhir bool = nilaiAkhir > 75
	var lulusNilaiAbsen bool = nilaiAbsen > 75

	var lulus bool = lulusNilaiAbsen || lulusNilaiAkhir
	fmt.Println(lulus)
}
