package main

import (
	"fmt"
)

func main() {
	counter := 1

	for counter <= 5 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	for counter := 1; counter <= 5; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	names := []string{"ivan", "satria"}

	for indeks, names := range names {
		fmt.Println("indeks ke", indeks, "=", names)
	}

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
}
