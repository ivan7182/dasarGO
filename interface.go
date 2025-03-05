package main

import "fmt"

type Animal interface {
	Speak() string
}

type Vehicle interface {
	Start() string
}

type Car struct{}
type Motorcycle struct{}

type Dog struct{}
type Cat struct{}

func (C Car) Start() string {
	return "Car is starting"
}

func (M Motorcycle) Start() string {
	return "Motorcycle is starting"
}

func Drive(vehicle Vehicle) {
	fmt.Println(vehicle.Start(), "Driving")
}

func Describe(describe interface{}) {
	fmt.Println("describe", describe)
}

func IntroduceAnimal(a Animal) {
	fmt.Println(a.Speak())
}

func (d Dog) Speak() string {
	return "Woof"
}

func (c Cat) Speak() string {
	return "meow"
}

func PrintValue(value interface{}) {
	fmt.Println("value", value)
}

func main() {

	var v Vehicle
	v = Car{}
	fmt.Println(v.Start())

	v = Motorcycle{}
	fmt.Println(v.Start())

	Drive(Car{})

	vehicles := []Vehicle{
		Car{},
		Motorcycle{},
		Car{},
	}

	for _, v := range vehicles {
		fmt.Println(v.Start())
	}

	Describe("Belajar Golang")
	Describe(2024)
	Describe(true)
	// var a Animal

	// a = Dog{}
	// fmt.Println(a.Speak())

	// a = Cat{}
	// fmt.Println(a.Speak())

	// IntroduceAnimal(Dog{})
	// PrintValue("Hello")
	// PrintValue(100)
	// PrintValue(true)
}
