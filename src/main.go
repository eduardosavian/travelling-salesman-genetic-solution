package main

import (
	"fmt"
	"strings"
)

func main() {
	array := readCSV("data/data.csv")

	var converted []uint64
	for _, n := range array {
		n, err := convertToUInt64(strings.TrimSpace(n))
		if err != nil {
			fmt.Println("Error converting element:", err)
			return
		}
		converted = append(converted, n)
	}

	size := converted[0]
	data := converted[1:]

	fmt.Println(size)
	fmt.Println(data)

	matrix := make([][]uint, size)
	for i := range matrix {
		matrix[i] = make([]uint, size)
	}

	// for i := 0; i < len(matrix); i++ {
	// 	fmt.Println(len(matrix))
	// 	for j := 0; j < len(matrix[i]); j++ {
	// 		//fmt.Printf("%d ", matrix[i][j])
	// 	}
	// 	fmt.Println()
	// }
}