package main

import (
	"fmt"
)

// Genetic algorithm function to calculate the best route and distance
func geneticAlgorithm(matrix [][]int, nGenerations int, populationSize int) ([]int, int) {
	printMatrix(matrix)
	fmt.Println("Number of generations:", nGenerations)
	fmt.Println("Population size:", populationSize)
	numCities := len(matrix)

	fmt.Println(generateRandomRoute(numCities))

	bestRoute := make([]int, numCities)
	bestDistance := 0

	return bestRoute, bestDistance
}
