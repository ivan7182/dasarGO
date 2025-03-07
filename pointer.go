package main

import "fmt"

type User struct {
	Name string
}

type Product struct {
	Name  string
	Stock int
}

func ChangeName(u *User, newName string) {
	u.Name = newName
}

func HalveValue(p *float64) {
	*p = *p / 2
}

func AddStock(ad *Product, newStock int) {
	ad.Stock = ad.Stock + newStock
}

func ToggleSign(value *int) {
	*value = -*value
}

func main() {
	// u := User{Name: "Andi"}
	// ChangeName(&u, "budi")
	// fmt.Println(u.Name)

	// x := 100.0
	// HalveValue(&x)
	// fmt.Println(x)

	// p := Product{Name: "Buku", Stock: 10}
	// AddStock(&p, 5)
	// fmt.Println(p.Stock)

	x := 7
	ToggleSign(&x)
	fmt.Println(x) // Output: -7

	y := -3
	ToggleSign(&y)
	fmt.Println(y)

}
