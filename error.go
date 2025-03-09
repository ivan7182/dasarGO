package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Pembagi nol")
	} else {
		return a / b, nil
	}
}

func main() {
	hasil, err := Divide(90, 0)
	if err == nil {
		fmt.Println("hasil", hasil)
	} else {
		fmt.Println("error", err.Error())
	}
}
