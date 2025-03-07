package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func MultiplyByTwo(mbt *int) {
	*mbt = *mbt * 2
}

func UpdatePrice(up *Product, harga *float64) {
	up.Price = *harga
}

func main() {
	// x := 5
	// MultiplyByTwo(&x)
	// fmt.Println(x)

	p := Product{Name: "Laptop", Price: 10000}
	newPrice := 12000.0
	UpdatePrice(&p, &newPrice)

	fmt.Println(p.Name, p.Price)

}
