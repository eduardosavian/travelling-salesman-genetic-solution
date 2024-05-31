package main

import "fmt"

func geneticAlgorithm(matrix [][]int, nGenerations int, populationSize int) ([]int, int) {
	printMatrix(matrix)
	fmt.Println("Number of generations:", nGenerations)
	fmt.Println("Population size:", populationSize)
	size := len(matrix)

	bestRoute := make([]int, size)
	bestDistance := 0

	return bestRoute, bestDistance
}
