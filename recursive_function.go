package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)

	}
}

func countDown(n int) {
	if n <= 0 {
		return
	}

	fmt.Println(n)

	countDown(n - 1)
}

func factorial(f int) int {
	if f == 0 || f == 1 {
		return 1
	}
	return f * factorial(f-1)
}

func main() {
	fmt.Println(factorialRecursive(10))
	countDown(5)
	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("Factorial of 3:", factorial(3))
	fmt.Println("Factorial of 1:", factorial(1))
	fmt.Println("Factorial of 0:", factorial(0))
}
