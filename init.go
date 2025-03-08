package main

import (
	"DASARGO/database"
	_ "DASARGO/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
