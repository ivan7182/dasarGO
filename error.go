package main

import (
	"errors"
	"fmt"
)

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	} else {
		return a / b, nil
	}
}

func main() {
	result, err := Divide(100, 10)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("hasil", result)
	}
}
