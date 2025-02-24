package main

import "fmt"

func main() {
	a := 10
	b := 20
	c := a + b

	fmt.Println(c)

	i := 20
	i += 10
	fmt.Println(i)

	i += 4
	fmt.Println(i)

	j := 20
	j++
	j++

	fmt.Println(j)

	j--
	j--
	fmt.Println(j)

}
