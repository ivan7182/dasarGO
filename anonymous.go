package main

import "fmt"

type Blaacklist func(name string) bool
type Validator func(string) bool

func registerUser(name string, blacklist Blaacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func validasiName(nama string, validator Validator) {
	if validator(nama) {
		fmt.Printf("nama %s valid\n", nama)
	} else {
		fmt.Printf("nama %s tidak valid\n", nama)
	}
}

func operasiKalkulator() (func(int, int) int, func(int, int) int) {

	kali := func(a, b int) int {
		return a * b
	}
	bagi := func(a, b int) int {
		if b == 0 {
			fmt.Println("error")
			return 0
		}
		return a / b
	}

	return kali, bagi
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("ivan", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

	validasiName("budi", func(nama string) bool {
		return len(nama) >= 5
	})

	validasiName("jonathan", func(nama string) bool {
		return len(nama) >= 5
	})

	kali, bagi := operasiKalkulator()

	fmt.Println("hasil", kali(4, 5))
	fmt.Println("hasil", bagi(10, 5))
}
