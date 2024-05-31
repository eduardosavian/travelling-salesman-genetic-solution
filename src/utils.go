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

func formatMatrix(data []int, size int) [][]int {
	maxUint := ^uint(0)
	maxInt := int(maxUint >> 1)

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
				matrix[i][j] = maxInt
			}
		}
	}

	return matrix
}

func printMatrix(matrix [][]int) {
	for _, array := range matrix {
		for j := range array {
			fmt.Print(array[j], " ")
		}
		fmt.Println()

	}
}
