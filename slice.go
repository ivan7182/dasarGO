package main

import "fmt"

func main() {
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[3:]

	daySlice1[0] = "kamis Baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "hari baru")
	fmt.Println(daySlice2)

	fmt.Println(days)

	slice1 := days[1:4]
	fmt.Println(len(slice1))

	var slice3 []string = days[4:6]
	fmt.Println(slice3)
	fmt.Println(cap(slice3))

	var slice4 []string = days[:]
	fmt.Println(slice4)

	//menambahkan slice
	data := []int{1, 2, 3}
	sliceData1 := append(data, 4)
	fmt.Println(sliceData1)

	//menggabungkan kedua slice
	slice1A := []string{"a", "b", "c"}
	slice2B := []string{"d", "e", "f"}

	sliceA := append(slice1A, slice2B...)
	fmt.Println(sliceA)

}
