package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Read CSV file function with the input row of operation like "4; 2; 5; 10; 8; 4; 7"
func readCSV(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(strings.NewReader(string(content)))
	reader.Comma = ';'

	matrix, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	records := matrix[0]

	return records
}

// Convert string to int function
func convertToInt(array []string) []int {
	var converted []int

	for _, n := range array {
		n, err := strconv.Atoi(strings.TrimSpace(n))
		if err != nil {
			panic("Error converting element:" + err.Error())
		}
		converted = append(converted, n)
	}

	return converted
}

// Format Matrix function from string of the CSV file
func formatMatrix(data []int, size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	dataIndex := 0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				matrix[i][j] = -1
			} else if i < j {
				matrix[i][j] = data[dataIndex]
				dataIndex++
			} else {
				matrix[i][j] = matrix[j][i]
			}
		}
	}

	return matrix
}

// Print Matrix function
func printMatrix(matrix [][]int) {
	fmt.Println("Matrix")
	for _, array := range matrix {
		for j := range array {
			fmt.Print(array[j], " ")
		}
		fmt.Println()
	}
}

// Calculate the total distance of a route
func calculateTotalDistance(matrix [][]int, route []int) int {
	totalDistance := 0
	for i := 0; i < len(route)-1; i++ {
		totalDistance += matrix[route[i]][route[i+1]]
	}
	totalDistance += matrix[route[len(route)-1]][route[0]]
	return totalDistance
}

// Find the best route in the population
func findBestRoute(matrix [][]int, population [][]int) ([]int, int) {
	bestRoute := make([]int, len(population[0]))
	bestDistance := int(^uint(0) >> 1)

	for _, route := range population {
		distance := calculateTotalDistance(matrix, route)
		if distance < bestDistance {
			bestDistance = distance
			copy(bestRoute, route)
		}
	}

	return bestRoute, bestDistance
}
