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
	fmt.Println("Matrix")
	for _, array := range matrix {
		for j := range array {
			fmt.Print(array[j], " ")
		}
		fmt.Println()
	}
}

