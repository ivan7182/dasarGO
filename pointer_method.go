package main

import "fmt"

type Man struct {
	Name string
}

type Car struct {
	Brand string
	Speed int
}

func (man *Man) Merried() {
	man.Name = "MR . " + man.Name
}

func MultiplyValue(p *int, factor int) {
	*p = *p * factor
}

func ResetToZero(p *int) {
	*p = 0
}

func (c *Car) IncreaseSpeed(amount int) {
	c.Speed += amount
}

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	// ivan := Man{"Ivan"}
	// ivan.Merried()

	// fmt.Println(ivan.Name)

	// x := 5
	// MultiplyValue(&x, 3)
	// fmt.Println(x)

	// xz := 100
	// ResetToZero(&xz)
	// fmt.Println(xz)

	a, b := 10, 20
	Swap(&a, &b)
	fmt.Println(a, b)

}
