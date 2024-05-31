package main

import (
	"fmt"
)

func main() {
	array := readCSV("data/data.csv")

	converted := convertToInt(array)

	size := converted[0]
	data := converted[1:]

	fmt.Println(size)
	fmt.Println(data)
	fmt.Println()

	if size < 2 || size > 20 {
		panic("Invalid graph size, must be gte 2 or lte 20")
	}

	if (size * size) / 2 - 2 != len(data) {
		panic("Invalid data size, must be equal to graph size")
	}

	matrix := formatMatrix(size, data)
	fmt.Println((matrix))
}