package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}

	return total
}

func cetakPesan(pesan ...string) {
	for _, pesan := range pesan {
		fmt.Println(pesan)
	}
}

func main() {
	fmt.Println(sumAll(10, 10, 10, 10, 10))

	number := []int{10, 10, 10, 10}
	fmt.Println(sumAll(number...))

	cetakPesan("halo", "ivan", "satria")
}
