package main

import "fmt"

func main() {
	names := []string{"ivan", "satria", "sari", "supri", "anwar"}

	var slice1 []string = names[1:4]
	fmt.Println(slice1)

	var slice []string = names[2:4]
	fmt.Println(slice)

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	var daySlice1 = days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "hari baru")
	daySlice2[0] = "sabtu lama"

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 3, 4)
	newSlice[0] = "muhammad"
	newSlice[2] = "satria"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "ivan")
	fmt.Println(newSlice2)

	newSlice2[0] = "budi"
	fmt.Println(newSlice2)

}
