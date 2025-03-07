package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Rectangle struct {
	Width, Height float64
}

func SquareValue(p *int) {
	*p = *p * *p
}

func NegateValue(nv *int) {
	*nv = -*nv
}

func IncreaseAge(p *Person) {
	p.Age++
}

func (r *Rectangle) Scale(factor float64) {
	r.Height = r.Height * factor
	r.Width = r.Width * factor
}

func main() {
	// x := 5
	// SquareValue(&x)
	// fmt.Println(x)

	x := 10
	NegateValue(&x)
	fmt.Println(x) // Output: -10

	y := -7
	NegateValue(&y)
	fmt.Println(y)

	p := Person{Name: "Budi", Age: 25}
	IncreaseAge(&p)
	fmt.Println(p.Age)

	r := Rectangle{Width: 4, Height: 5}
	r.Scale(2)
	fmt.Println(r.Width, r.Height)
}
