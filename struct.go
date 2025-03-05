package main

import (
	"fmt"
)

type Customer struct {
	Name, Address string
	Age           int
}

type Mahasiswa struct {
	Nama, NIM string
	IPK       float64
}

func (customer Customer) sayHello(name string) {
	fmt.Println("hello", name, "my name is", customer.Name, "umur saya", customer.Age)
}

func main() {
	var ivan Customer
	ivan.Name = "Ivan"
	ivan.Address = "Tegal"
	ivan.Age = 24
	fmt.Println(ivan)
	fmt.Println(ivan.Name)

	joko := Customer{
		Name:    "joko",
		Address: "tegal",
		Age:     22,
	}
	fmt.Println(joko)

	budi := Customer{"budi", "jakarta", 22}
	fmt.Println(budi)

	satria := Mahasiswa{
		Nama: "staria",
		NIM:  "19090082",
		IPK:  14.4,
	}
	fmt.Println(satria)

	joko.sayHello("ivan")

}
