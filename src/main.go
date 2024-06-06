package main

import (
	"fmt"
)

func main() {
	array := readCSV("data/data.csv")

	converted := convertToInt(array)

	size := converted[0]
	data := converted[1:]

	fmt.Println("Graph size:", size)
	fmt.Println("Data:", data)

	if size < 2 || size > 20 {
		panic("Invalid graph size, must be gte 2 or lte 20")
	}

	if (size*size) / 2 - 2 != len(data) {
		panic("Invalid data size, must be equal to graph size")
	}

	matrix := formatMatrix(data, size)

	numGenerations := 100
	populationSize := 50

	printMatrix(matrix)
	fmt.Println("Number of generations:", numGenerations)
	fmt.Println("Population size:", populationSize)

	bestRoute, bestDistance := geneticAlgorithm(matrix, numGenerations, populationSize)

	fmt.Println("Best Route:", bestRoute)
	fmt.Println("Best Distance:", bestDistance)
}
