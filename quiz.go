package main

import (
	"fmt"
	"strings"
)

type UsernameError struct {
	Message string
}

type AgeError struct {
	Message string
}

func (a *AgeError) Error() string {
	return a.Message
}

func ValidateAge(age int) error {
	if age < 18 {
		return &AgeError{Message: "Usia minimal 18 tahun"}
	}
	if age > 150 {
		return &AgeError{Message: "Usia tidak realistis"}
	}
	return nil
}

func (u *UsernameError) Error() string {
	return u.Message
}

func ValidateUsername(nama string) error {
	if len(nama) < 5 {
		return &UsernameError{Message: "Username terlalu pendek"}
	}

	if strings.Contains(nama, " ") {
		return &UsernameError{Message: "Username tidak boleh mengandung spasi"}
	}

	return nil
}

func main() {
	// err := ValidateUsername("unaan")
	// if err != nil {
	// 	fmt.Println(err) // Output: Username terlalu pendek
	// } else {
	// 	fmt.Println("sukses")
	// }

	err := ValidateAge(160)
	if err != nil {
		fmt.Println(err) // Output: Usia minimal 18 tahun
	} else {
		fmt.Println("sukses")
	}

}
