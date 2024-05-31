package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
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

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			matrix[i][j] = data[i+j]
		}
	}

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			matrix[j][i] = data[i+j]
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				matrix[i][j] = -1
			}
		}
	}

	return matrix
}

// Print Matrix function
func printMatrix(matrix [][]int) {
	for _, array := range matrix {
		for j := range array {
			fmt.Print(array[j], " ")
		}
		fmt.Println()
	}
}

func contains(route []int, city int) bool {
	for _, c := range route {
		if c == city {
			return true
		}
	}
	return false
}

func generateRandomRoute(numCities int) []int {
	route := make([]int, numCities)
	for i := range route {
		route[i] = i
	}
	rand.Shuffle(len(route), func(i, j int) { route[i], route[j] = route[j], route[i] })
	return route
}

func distance(matrix [][]int, city1 int, city2 int) int {
	if city1 == city2 {
		return 0
	}
	if matrix[city1][city2] == -1 {
		return math.MaxInt32
	}
	return matrix[city1][city2]
}

func totalDistance(matrix [][]int, route []int) int {
	total := 0
	for i := 0; i < len(route)-1; i++ {
		total += distance(matrix, route[i], route[i+1])
	}
	total += distance(matrix, route[len(route)-1], route[0])
	return total
}
