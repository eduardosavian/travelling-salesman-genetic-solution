package main

import (
	"fmt"
)

func main() {
	array := readCSV("data/data.csv")
	size := array[0]
	data := array[1:]

	fmt.Println(size)
	fmt.Println(data)

}