package main

import "fmt"

func main() {
	var nilai = [4]int{90, 80, 60, 50}

	fmt.Println(nilai)

	var nilaiAbsen = [...]int{
		50,
		80,
		60,
		70,
	}

	fmt.Println(len(nilaiAbsen))
}
